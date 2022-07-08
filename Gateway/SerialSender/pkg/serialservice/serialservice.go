package serialservice

import (
	serialreceiver "SerialSender/pkg/serialservice/serialLogic"
	serialAccess "SerialSender/pkg/serialservice/serialaccess"
	"context"
)

type iSerialService interface {
	Send(ctx context.Context, dataBuffer []byte)
	Close() error
}

type SerialService struct {
	service iSerialService
}

func ServiceServiceFactory(portToOpen string) *SerialService {
	return &SerialService{
		service: serialreceiver.FactorySerialReciever(serialAccess.NewSerialPortReader(portToOpen)),
	}
}

func (s *SerialService) Send(ctx context.Context, dataBuffer []byte) {
	s.service.Send(ctx, dataBuffer)
}
func (s *SerialService) Close() error {
	return s.service.Close()
}
