package serialreciever

import (
	serialAccess "SerialSender/pkg/serialservice/serialaccess"
	"context"
)

type SerialReciever struct {
	serialAbs serialAccess.ISerialConnection
}

func FactorySerialReciever(serialAbs serialAccess.ISerialConnection) *SerialReciever {
	return &SerialReciever{
		serialAbs: serialAbs,
	}
}

func (s *SerialReciever) Send(ctx context.Context, dataBuffer []byte) {
	s.serialAbs.Send(ctx, dataBuffer)
}

func (s *SerialReciever) Close() error {
	return s.serialAbs.Close()
}
