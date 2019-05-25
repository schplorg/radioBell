package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type Server struct {
	srv       *http.Server
	clients   map[*Client]bool
	broadcast chan []byte
	files     []string
}
type Client struct {
	send chan []byte
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (serv *Server) stopServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := serv.srv.Shutdown(ctx)
	if err != nil {
		fmt.Printf("Shutdown failed: %s", err)
	}
	return err
}

func createServer() *Server {
	mux := http.NewServeMux()
	srv := &http.Server{Addr: ":58000", Handler: mux}
	serv := Server{
		srv:       srv,
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte),
	}
	// check for new messages to broadcast
	go func() {
		defer func() {
			fmt.Println("ending http serve")
		}()
		fs := http.FileServer(http.Dir("public"))
		mux.Handle("/", fs)

		fmt.Println("serv.serveWs")
		mux.HandleFunc("/ws", serv.serveWs)
		err := serv.srv.ListenAndServe()
		if err != nil {
			fmt.Println("ListenAndServe error!")
		}
	}()
	go func() {
		defer func() {
			fmt.Println("ending broadcast")
		}()
		fmt.Println("serv.broadcast")
		ticker := time.NewTicker(time.Second)
		defer func() {
			fmt.Println("ending ticker")
			ticker.Stop()
		}()
		for {
			select {
			case message := <-serv.broadcast:
				for c := range serv.clients {
					c.send <- message
				}
			case <-ticker.C:
				fmt.Println("tick")
				for c := range serv.clients {
					c.send <- []byte("tick")
				}
			}
		}
	}()
	return &serv
}

func (serv *Server) serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serveWs")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error!")
	}
	client := Client{send: make(chan []byte, 256), conn: conn}
	serv.clients[&client] = true
	// check for new messages to send
	go func() {
		defer func() {
			fmt.Println("ending websocket")
			client.conn.Close()
			delete(serv.clients, &client)
		}()
		for {
			select {
			case message := <-client.send:
				{
					writer, err := client.conn.NextWriter(websocket.TextMessage)
					if err != nil {
						fmt.Println("NextWriter error!")
						return
					}
					writer.Write(message)
					n := len(client.send)
					for i := 0; i < n; i++ {
						_, err = writer.Write(<-client.send)
						if err != nil {
							fmt.Println("Write error!")
							return
						}
					}
					err = writer.Close()
					if err != nil {
						fmt.Println("Write error!")
						return
					}
				}
			}
		}
	}()
}
