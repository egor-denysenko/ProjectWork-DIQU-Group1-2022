package serialservice

import (
	"context"
	serialAccess "serialReciever/pkg/serialaccess"
	serialReciever "serialReciever/pkg/serialreciever"
)

type iSerialService interface {
	Recieve(ctx context.Context, out chan<- []byte)
	Close() error
}

type SerialService struct {
	service iSerialService
}

func ServiceServiceFactory(portToOpen string) *SerialService {
	return &SerialService{
		service: serialReciever.FactorySerialReciever(serialAccess.NewSerialPortReader(portToOpen)),
	}
}

func (s *SerialService) Recieve(ctx context.Context, out chan<- []byte) {
	s.service.Recieve(ctx, out)
}
func (s *SerialService) Close() error {
	return s.service.Close()
}
