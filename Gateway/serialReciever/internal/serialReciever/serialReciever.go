package serialreciever

import (
	"fmt"
	"log"
	"math/bits"

	"go.bug.st/serial"
)

const baudRate = 9600
const numberOfDataBits = 8
const bytesContainingInformation int = 4

type SerialReciever interface {
	PortOpener(portToOpen string) serial.Port
	RecieveSerial(portConnection serial.Port) []byte
}

type SerialRecieverImplementation struct{}

func (s SerialRecieverImplementation) PortOpener(portToOpen string) serial.Port {
	port, err := serial.Open(portToOpen, &serial.Mode{
		BaudRate: baudRate,
		Parity:   serial.NoParity,
		DataBits: numberOfDataBits,
		StopBits: serial.OneStopBit,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reading on port: COM6")
	return port
}

func (s SerialRecieverImplementation) RecieveSerial(portConnection serial.Port) []byte {
	buff := make([]byte, bytesContainingInformation)
	fmt.Println("cerco")
	n, err := portConnection.Read(buff)
	fmt.Println(n)
	if err != nil {
		log.Fatal(err)
	}
	if n == 0 {
		fmt.Println("\nEOF")
		//	break
	}
	fmt.Println("bits in 0 buffer", bits.OnesCount(uint(buff[0])))
	fmt.Printf("Buffer Recieved in index 0 %b \n", buff[0])
	fmt.Println("Buffer Recieved ", buff)
	fmt.Println("Buffer Recieved ", buff[:n])
	return buff
}

/*func serialPortsRetriever() []string {
	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	// Print the list of detected ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
	return ports
}*/

func main() {
	test := SerialRecieverImplementation{}
	test.RecieveSerial(test.PortOpener("COM6"))

	// in quache modo ora devo parsare questo dato in modo da usarlo in qualche modo zio kentaro...
	// non posso associare a ciascun valore una roba da scrivere con una serie di if è bruttissimo
}
