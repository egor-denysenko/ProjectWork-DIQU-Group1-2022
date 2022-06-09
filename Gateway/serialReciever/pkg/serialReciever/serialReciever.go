package serialReciever

import serialrecieverabs "serialReciever/pkg/serialReciever/serialRecieverAbstraction"

type iserialReciever interface {
	Recieve() (data []byte, err error)
	Close() error
}

type serialReciever struct {
	serialAbs serialrecieverabs.ISerialPortReader
}

func FactorySerialReciever(serialAbs serialrecieverabs.ISerialPortReader) *serialReciever {
	return &serialReciever{
		serialAbs: serialAbs,
	}
}

func (s *serialReciever) Recieve() ([]byte, error) {
	return s.Recieve()
}

func (s *serialReciever) Close() error {
	return s.Close()
}
