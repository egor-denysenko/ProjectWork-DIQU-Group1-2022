package serialreciever

import serialAccess "serialReciever/pkg/serialaccess"

type SerialReciever struct {
	serialAbs serialAccess.ISerialConnection
}

func FactorySerialReciever(serialAbs serialAccess.ISerialConnection) *SerialReciever {
	return &SerialReciever{
		serialAbs: serialAbs,
	}
}

func (s *SerialReciever) Recieve() ([]byte, error) {
	return s.serialAbs.Recieve()
}

func (s *SerialReciever) Close() error {
	return s.serialAbs.Close()
}
