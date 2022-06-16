package protocolParser

import (
	"context"
	"errors"
)

type RecieveCommand uint8

const (
	RecieveData RecieveCommand = 69
)

var MessageNotForTheGateway = errors.New("message not for the gateway")

type FormattedData struct {
	locomotiveID uint
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
func ParseSerialData(ctx context.Context, recievedSerial []byte) error {
	errReciever := DetermineReciever(recievedSerial[1])
	if errReciever != nil {
		return errReciever
	}

	//commandValidity, commandToExcecute := DetermineCommand(recievedSerial[2])

	var StagingDataStruct FormattedData

	StagingDataStruct.deviceID = recievedSerial[0]
	StagingDataStruct.locomotiveID = 188
	//testChannel <- json.Marshal(StagingDataStruct)
	return nil
}

func DetermineReciever(recieverByte uint8) error {
	var gatewayId uint8 = 254
	if recieverByte != gatewayId {
		return MessageNotForTheGateway
	}
	return nil
}

func DetermineCommand(commandByte uint8) (bool, RecieveCommand) {
	var parsedCommand = RecieveCommand(commandByte)
	switch parsedCommand {
	case RecieveData:
		return true, parsedCommand
	}
	return false, 0
}
