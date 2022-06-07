package serialreciever

import (
	"fmt"
	"log"
	"math/bits"

	"go.bug.st/serial"
)

type iSerialPortReader interface {
	Read() (data []byte, err error)
	Close() error
}

type SerialConnection struct {
	iSerialPortReader    iSerialPortReader
	serialPortConnection serial.Port
}

func NewSerialPortReader() SerialConnection {
	port, err := serial.Open("COM7", &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   0,
		StopBits: 0,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reading on port: COM6")
	return SerialConnection{
		serialPortConnection: port,
	}
}

func (s *SerialConnection) RecieveSerial() ([]byte, error) {
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
	fmt.Println("bits in 0 buffer", bits.OnesCount(uint(buff[0])))
	fmt.Printf("Buffer Recieved in index 0 %b \n", buff[0])
	fmt.Println("Buffer Recieved ", buff)
	fmt.Println("Buffer Recieved ", buff[:n])
	return buff, nil
}
