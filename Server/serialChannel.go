package main

import (
	"fmt"
	"github.com/tarm/serial"
)

type SerialChannel struct {
	serialChannel chan string
}

func createSerialChannel() *SerialChannel {
	sc := &SerialChannel{
		serialChannel: make(chan string),
	}
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("fail")
	} else {
		go sc.readSerial(s)
	}
	return sc
}
func (sc *SerialChannel) readSerial(s *serial.Port) {
	for true {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			fmt.Println("fail")
		}
		out := string(buf[:n])
		sc.serialChannel <- out
	}
}
