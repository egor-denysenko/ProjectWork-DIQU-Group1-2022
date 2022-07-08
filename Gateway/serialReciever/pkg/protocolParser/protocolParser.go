package protocolParser

import (
	"encoding/json"
	"errors"
	"log"
)

type RecieveCommand uint8

const (
	RecieveData RecieveCommand = 69
)

var MessageNotForTheGateway = errors.New("message not for the gateway")
var WrongGatewayCommand = errors.New("the command doesn't exist")

type FormattedData struct {
	LocomotiveID uint
	VagonID      uint8
	Temperature  uint8
	Humidity     uint8
	VagonAllarms VagonAllarms
	VagonDoors   VagonDoors
	VagonLights  VagonLights
}

type VagonAllarms struct {
	DoorIO         bool
	DoorB          bool
	DoorC          bool
	TemperatoreMin bool
	TemperatureMax bool
	Lights         bool
	Humidity       bool
}

type VagonDoors struct {
	Door1       bool
	Door2       bool
	Door3       bool
	Door4       bool
	DoorBath    bool
	DoorConduct bool
}

type VagonLights struct {
	LightMode   bool
	LightStatus bool
}

func ValidateSerialData(serialData []byte, parsedDataChan chan<- []byte) {
	errReciever := determineReciever(serialData[0])
	log.Printf("errReciever %v", errReciever)
	if errReciever != nil {
		parsedDataChan <- nil
		return
	}
	errCommand, recievedCommand := determineCommand(serialData[2])
	log.Printf("errCommand %v", errCommand)
	if errCommand == WrongGatewayCommand {
		parsedDataChan <- nil
		return
	}
	switch recievedCommand {
	case RecieveData:
		parsedDataChan <- parseSerialData(serialData)
		return
	}
}

//Parse known bytes in the FromattedData struct and call other functions that will analyze each bits for specific bytes
func parseSerialData(recievedSerial []byte) []byte {
	StagingDataStruct := FormattedData{}
	//var StagingDataStruct FormattedData

	StagingDataStruct.VagonID = recievedSerial[1]
	StagingDataStruct.LocomotiveID = 188
	StagingDataStruct.Temperature = 30
	StagingDataStruct.Humidity = 80
	StagingDataStruct.VagonAllarms = VagonAllarms{
		DoorIO:         false,
		DoorB:          false,
		DoorC:          false,
		TemperatoreMin: false,
		TemperatureMax: false,
		Lights:         false,
		Humidity:       false,
	}
	StagingDataStruct.VagonDoors = VagonDoors{
		Door1:       false,
		Door2:       false,
		Door3:       false,
		Door4:       false,
		DoorBath:    false,
		DoorConduct: false,
	}
	StagingDataStruct.VagonLights = VagonLights{
		LightMode:   false,
		LightStatus: false,
	}
	jsonData, err := json.Marshal(StagingDataStruct)
	if err != nil {
		log.Fatalln(err)
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
