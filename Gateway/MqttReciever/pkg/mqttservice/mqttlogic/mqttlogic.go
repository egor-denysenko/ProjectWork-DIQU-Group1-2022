package mqttlogic

import "context"

type IMqttLogic interface {
	Connect() error
	Subscribe(ctx context.Context, topic string) error
}

type MqttPubblishOptions struct {
}

type MqttLogic struct {
	mqttAbs IMqttLogic
}

func FactoryMqttLogic(mqttAccess IMqttLogic) *MqttLogic {
	return &MqttLogic{mqttAbs: mqttAccess}
}

func (m *MqttLogic) Connect() error {
	return m.mqttAbs.Connect()
}
func (m *MqttLogic) Subscribe(ctx context.Context, topic string) error {
	return m.mqttAbs.Subscribe(ctx, topic)
}
