package main

import (
	"fmt"
	"github.com/tarm/serial"
	"time"
)

func main() {
	fmt.Println("serial")
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("fail")
	} else {
		go readSerial(s)
	}
	loop()
}

func loop() {
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}
func readSerial(s *serial.Port) {
	for true {
		buf := make([]byte, 128)
		fmt.Println("read")
		n, err := s.Read(buf)
		if err != nil {
			fmt.Println("fail")
		}
		out := string(buf[:n])
		fmt.Println(out)
	}
}
