package protocolParser

import (
	"context"
	"encoding/json"
	"errors"
)

var MessageNotForTheGateway = errors.New("message not for the gateway")

type FormattedData struct {
	locomotiveID uint8
	deviceID     uint8
	temperature  uint8
	humidity     uint8
	vagonStatus  VagonStatus
	vagonAlarms  VagonAlarms
}

type VagonAlarms struct {
}
type VagonStatus struct {
	withBathroom bool
}

//Parse known bytes in the FromattedData struct and call other functions that will analyze each bits for specific bytes
func ParseSerialData(ctx context.Context, recievedSerial []byte) ([]byte, error) {

	var StagingDataStruct FormattedData

	return json.Marshal(StagingDataStruct)
}

func DetermineReciever(recieverByte uint8) error {
	var gatewayId uint8 = 254
	if recieverByte != gatewayId {
		return MessageNotForTheGateway
	}
	return nil
}
