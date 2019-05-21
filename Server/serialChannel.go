package main

import (
	"fmt"
	"github.com/tarm/serial"
	"os"
)

type SerialChannel struct {
	receive chan string
}

func createSerialChannel() *SerialChannel {

	ports := func() (ports []string) {
		ports = make([]string, 0)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered from:", err)
				ports = []string{
					"/dev/ttyUSB0",
					"/dev/ttyUSB1",
					"COM8",
				}
			}
		}()
		if len(os.Args) <= 1 {
			panic("too few args!")
		}
		argi := os.Args[1]
		ports = append(ports, argi)
		return
	}()

	sc := &SerialChannel{
		receive: make(chan string),
	}
	for _, p := range ports {
		c := &serial.Config{Name: p, Baud: 9600}
		s, err := serial.OpenPort(c)
		if err != nil {
			fmt.Println("OpenPort " + p + " error!")
		} else {
			go sc.readSerial(s)
			break
		}
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
