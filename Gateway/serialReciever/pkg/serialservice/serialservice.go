package serialservice

import (
	"context"
	"serialReciever/pkg/serialservice/serialReciever"
	serialAccess "serialReciever/pkg/serialservice/serialaccess"
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
		service: serialreciever.FactorySerialReciever(serialAccess.NewSerialPortReader(portToOpen)),
	}
}

func (s *SerialService) Recieve(ctx context.Context, out chan<- []byte) {
	s.service.Recieve(ctx, out)
}
func (s *SerialService) Close() error {
	return s.service.Close()
}
