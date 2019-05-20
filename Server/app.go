package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	defer func() {
		fmt.Println("ending app")
	}()
	fmt.Println("creating server")
	serv := createServer()
	fmt.Println("creating serial")
	ser := createSerialChannel()
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer func() {
			fmt.Println("ending app ticker")
			ticker.Stop()
		}()
		beat := true
		for {
			select {
			case <-ticker.C:
				{
					if !beat {
						fmt.Println("heartbeat timeout!!")
						serv.srv.Shutdown(context.TODO())
						time.Sleep(2 * time.Second)
						serv = createServer()
						time.Sleep(2 * time.Second)
					} else {
						beat = false
					}
				}
			case msg := <-serv.heartBeat:
				{
					fmt.Println("hearbeat: " + msg)
					beat = true
				}
			}
		}
	}()
	// check for serial messages and broadcast them
	for {
		select {
		case message := <-ser.receive:
			fmt.Print(message)
			serv.broadcast <- []byte(message)
		}
	}
}
