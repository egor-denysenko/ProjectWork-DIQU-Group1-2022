package mqttlogic

type IMqttLogic interface {
	Connect() error
	Pubblish(topic string, message []byte, qos uint8) error
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
func (m *MqttLogic) Pubblish(topic string, message []byte, qos uint8) error {
	return m.mqttAbs.Pubblish(topic, message, qos)
}
