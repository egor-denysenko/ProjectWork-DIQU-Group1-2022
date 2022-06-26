package serialconfig

type SerialConfig struct {
	BaudRate int
	DataBits int
}

func NewSerialConfig() *SerialConfig {
	return &SerialConfig{
		BaudRate: 9600,
		DataBits: 8,
	}
}
