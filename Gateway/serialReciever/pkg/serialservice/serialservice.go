package serialservice

import (
	serialAccess "serialReciever/pkg/serialaccess"
	serialReciever "serialReciever/pkg/serialreciever"
)

type iSerialService interface {
	Recieve() (data []byte, err error)
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

func (s *SerialService) Recieve() ([]byte, error) {
	return s.service.Recieve()
}
func (s *SerialService) Close() error {
	return s.service.Close()
}
