package serialaccess

import (
	"fmt"
	"go.bug.st/serial"
	"log"
)

type ISerialConnection interface {
	Recieve() (data []byte, err error)
	Close() error
}

type SerialConnection struct {
	serialPortConnection serial.Port
}

func NewSerialPortReader(portToOpen string) *SerialConnection {
	port, err := serial.Open(portToOpen, &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   0,
		StopBits: 0,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reading on port:", portToOpen)
	return &SerialConnection{
		serialPortConnection: port,
	}
}

func (s *SerialConnection) Recieve() ([]byte, error) {
	buff := make([]byte, 4)
	fmt.Println("cerco")
	n, err := s.serialPortConnection.Read(buff)
	fmt.Println(n)
	if err != nil {
		log.Fatal(err)
		return buff, err
	}
	if n == 0 {
		fmt.Println("\nEOF")
		//	break
	}
	return buff, nil
}

func (s *SerialConnection) Close() error {
	return s.serialPortConnection.Close()
}
