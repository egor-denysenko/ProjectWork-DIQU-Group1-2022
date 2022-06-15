package protocolParser

import (
	"encoding/json"
)

type FormattedData struct {
	trainID     uint8
	deviceID    uint8
	temperature uint8
	humidity    uint8
	vagonStatus VagonStatus
}

type VagonStatus struct {
	withBathroom bool
}

//Parse known bytes in the FromattedData struct and call other functions that will analyze each bits for specific bytes
func ParseSerialData(recievedSerial []byte) ([]byte, error) {

	var StagingDataStruct FormattedData

	return json.Marshal(StagingDataStruct)
}
