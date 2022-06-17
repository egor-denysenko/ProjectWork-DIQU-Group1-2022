package protocolParser

import (
	"context"
	"encoding/json"
	"errors"
)

type RecieveCommand uint8

const (
	RecieveData RecieveCommand = 69
)

var MessageNotForTheGateway = errors.New("message not for the gateway")
var WrongGatewayCommand = errors.New("the command doesn't exist")

type FormattedData struct {
	locomotiveID uint
	vagonID      uint8
	temperature  uint8
	humidity     uint8
	vagonAllarms vagonAllarms
	vagonDoors   vagonDoors
	vagonLights  vagonLights
}

type vagonAllarms struct {
	doorIO         bool
	doorB          bool
	doorC          bool
	temperatoreMin bool
	temperatureMax bool
	lights         bool
	humidity       bool
}

type vagonDoors struct {
	door1       bool
	door2       bool
	door3       bool
	door4       bool
	doorBath    bool
	doorConduct bool
}

type vagonLights struct {
	lightMode   bool
	lightStatus bool
}

func ValidateSerialData(ctx context.Context, serialDataChan <-chan []byte, parsedData chan<- []byte) {
	recievedSerial := <-serialDataChan
	errReciever := determineReciever(recievedSerial[1])
	if errReciever != nil {
		parsedData <- nil
	}
	errCommand, recievedCommand := determineCommand(recievedSerial[2])
	if errCommand == WrongGatewayCommand {
		parsedData <- nil
	}
	switch recievedCommand {
	case RecieveData:
		parsedData <- parseSerialData(ctx, recievedSerial)
	default:
		parsedData <- nil
	}
}

//Parse known bytes in the FromattedData struct and call other functions that will analyze each bits for specific bytes
func parseSerialData(ctx context.Context, recievedSerial []byte) []byte {
	var StagingDataStruct FormattedData

	StagingDataStruct.vagonID = recievedSerial[0]
	StagingDataStruct.locomotiveID = 188
	StagingDataStruct.temperature = 30
	StagingDataStruct.humidity = 80
	jsonData, err := json.Marshal(StagingDataStruct)
	if err != nil {
		return nil
	}
	return jsonData
}

func determineReciever(recieverByte uint8) error {
	var gatewayId uint8 = 254
	if recieverByte != gatewayId {
		return MessageNotForTheGateway
	}
	return nil
}

func determineCommand(commandByte uint8) (error, RecieveCommand) {
	var parsedCommand = RecieveCommand(commandByte)
	switch parsedCommand {
	case RecieveData:
		return nil, parsedCommand
	}
	return WrongGatewayCommand, 0
}
