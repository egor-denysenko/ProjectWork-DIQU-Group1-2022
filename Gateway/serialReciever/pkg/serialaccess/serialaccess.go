package serialaccess

import (
	"context"
	"fmt"
	"log"

	"go.bug.st/serial"
)

type ISerialConnection interface {
	Recieve(ctx context.Context, out chan<- []byte)
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

func (s *SerialConnection) Recieve(ctx context.Context, out chan<- []byte) {
	buff := make([]byte, 7)
	for {
		n, err := s.serialPortConnection.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("%v", string(buff[:n]))
		out <- buff
	}
}
func (s *SerialConnection) Close() error {
	return s.serialPortConnection.Close()
}
