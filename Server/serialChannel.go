package main

import (
	"fmt"

	"github.com/tarm/serial"
)

type SerialChannel struct {
	receive chan string
}

func createSerialChannel() *SerialChannel {
	sc := &SerialChannel{
		receive: make(chan string),
	}
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("OpenPort /dev/ttyUSB0 error!")
		c := &serial.Config{Name: "/dev/ttyUSB1", Baud: 9600}
		s, err := serial.OpenPort(c)
		if err != nil {
			fmt.Println("OpenPort /dev/ttyUSB1 error!")
		} else {
			go sc.readSerial(s)
		}
	} else {
		go sc.readSerial(s)
	}
	return sc
}
func (sc *SerialChannel) readSerial(s *serial.Port) {
	defer func() {
		fmt.Println("ending read serial")
	}()
	for true {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			fmt.Println("Read error!")
		}
		out := string(buf[:n])
		sc.receive <- out
	}
}
