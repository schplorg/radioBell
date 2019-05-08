package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type Server struct {
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

func createServer() *Server {
	serv := Server{
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte),
		files:     []string{"index.html", "index.js"},
	}
	// check for new messages to broadcast
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "index.html")
		})
		http.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "index.js")
		})
		http.HandleFunc("/dmrdrn2.mp3", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "dmrdrn2.mp3")
		})
		fmt.Println("serv.serveWs")
		http.HandleFunc("/ws", serv.serveWs)
		http.ListenAndServe(":58000", nil)
	}()
	go func() {
		fmt.Println("serv.broadcast")
		ticker := time.NewTicker(time.Second)
		defer func() {
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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error!")
	}
	client := Client{send: make(chan []byte, 256), conn: conn}
	serv.clients[&client] = true
	// check for new messages to send
	go func() {
		defer func() {
			client.conn.Close()
		}()
		for {
			select {
			case message := <-client.send:
				{
					writer, err := client.conn.NextWriter(websocket.TextMessage)
					if err != nil {
						fmt.Println("error!")
						return
					}
					writer.Write(message)
					n := len(client.send)
					for i := 0; i < n; i++ {
						writer.Write(<-client.send)
					}
					writer.Close()
				}
			}
		}
	}()
}
