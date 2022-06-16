package protocolParser

import (
	"context"
	"errors"
	"fmt"
)

type RecieveCommand uint8

const (
	RecieveData RecieveCommand = 69
)

var MessageNotForTheGateway = errors.New("message not for the gateway")
var WrongGatewayCommand = errors.New("the command doesn't exist")

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
	errCommand, recievedCommand := DetermineCommand(recievedSerial[2])
	if errCommand == WrongGatewayCommand {
		return WrongGatewayCommand
	}
	fmt.Println(recievedCommand)
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

func DetermineCommand(commandByte uint8) (error, RecieveCommand) {
	var parsedCommand = RecieveCommand(commandByte)
	switch parsedCommand {
	case RecieveData:
		return nil, parsedCommand
	}
	return WrongGatewayCommand, 0
}
