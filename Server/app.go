package main

import (
	"fmt"
	// "time"
)

func main() {
	sc := createSerialChannel()
	for {
		select {
		case s := <-sc.serialChannel:
			fmt.Println(s)
		}
	}
}

// func loop() {
// 	ticker := time.NewTicker(time.Second)
// 	for t := range ticker.C {
// 		fmt.Println("Tick at", t)
// 	}
// }
