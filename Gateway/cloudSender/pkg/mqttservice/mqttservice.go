package mqttservice

import (
	"cloudSender/pkg/mqttservice/mqttaccess"
	"cloudSender/pkg/mqttservice/mqttlogic"
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
func (m *MqttService) Pubblish(topic, message string, qos uint8) error {
	return m.service.Pubblish(topic, message, qos)

}
