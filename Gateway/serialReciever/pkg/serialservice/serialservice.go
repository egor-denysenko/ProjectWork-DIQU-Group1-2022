package serialservice

import (
	serialAccess "serialReciever/pkg/serialaccess"
	serialReciever "serialReciever/pkg/serialreciever"
)

type iSerialService interface {
	Recieve() (data []byte, err error)
	Close() error
}

type serialService struct {
	service iSerialService
}

func ServiceServiceFactory() *serialService {
	return &serialService{
		service: serialReciever.FactorySerialReciever(serialAccess.NewSerialPortReader()),
	}
}

func (s *serialService) Recieve() ([]byte, error) {
	return s.service.Recieve()
}
func (s *serialService) Close() error {
	return s.service.Close()
}
