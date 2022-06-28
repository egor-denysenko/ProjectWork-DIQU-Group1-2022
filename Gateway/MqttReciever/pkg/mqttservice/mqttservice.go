package mqttservice

import (
	"MqttReciever/pkg/mqttservice/mqttaccess"
	"MqttReciever/pkg/mqttservice/mqttlogic"
	"context"
)

type MqttService struct {
	service *mqttlogic.MqttLogic
}

func FactoryMqttService() *MqttService {
	return &MqttService{service: mqttlogic.FactoryMqttLogic(mqttaccess.NewMqttConnection())}
}

func (m *MqttService) Connect() error {
	return m.service.Connect()
}
func (m *MqttService) Subscribe(ctx context.Context, topic string) error {
	return m.service.Subscribe(ctx, topic)
}
