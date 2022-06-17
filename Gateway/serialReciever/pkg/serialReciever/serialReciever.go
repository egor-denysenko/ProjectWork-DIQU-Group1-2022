package serialreciever

import (
	"golang.org/x/net/context"
	serialAccess "serialReciever/pkg/serialaccess"
)

type SerialReciever struct {
	serialAbs serialAccess.ISerialConnection
}

func FactorySerialReciever(serialAbs serialAccess.ISerialConnection) *SerialReciever {
	return &SerialReciever{
		serialAbs: serialAbs,
	}
}

func (s *SerialReciever) Recieve(ctx context.Context, out chan<- []byte) {
	s.serialAbs.Recieve(ctx, out)
}

func (s *SerialReciever) Close() error {
	return s.serialAbs.Close()
}
