package protocolParser

type FormattedData struct {
	deviceID    uint
	temperature uint8
	humidity    uint8
	vagonStatus VagonStatus
}

type VagonStatus struct {
	withBathroom bool
}

func ParseSerialData([]byte) []byte {
	return []byte{0}
}
