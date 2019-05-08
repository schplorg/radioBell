package main

import (
	"fmt"
)

func main() {
	fmt.Println("creating server")
	serv := createServer()
	fmt.Println("creating serial")
	ser := createSerialChannel()
	// check for serial messages and broadcast them
	for {
		select {
		case message := <-ser.receive:
			fmt.Print(message)
			serv.broadcast <- []byte(message)
		}
	}
}
