package protocolParser

import (
	"encoding/json"
	"errors"
	"fmt"
)

type RecieveCommand uint8

const (
	SetTemp  RecieveCommand = 70
	SetLight RecieveCommand = 71
)

var WrongGatewayCommand = errors.New("the command doesn't exist")

func ParseMessageToByte(messageToParse *string, parsedDataChan chan<- []byte) {
	data, err := messageToMap(messageToParse)
	if err != nil {
		parsedDataChan <- nil
	}
	errCommand, command := determineCommand(data["WagonCommand"])
	if errCommand != nil {
		parsedDataChan <- nil
	}
	switch command {
	case SetTemp:
		parsedDataChan <- CreateSetTempMess(data)
	case SetLight:
		parsedDataChan <- CreateSetLightMess(data)
	}
	parsedDataChan <- []byte{42}
}

func messageToMap(messageToParse *string) (map[string]uint8, error) {
	var usableData = make(map[string]uint8)
	err := json.Unmarshal([]byte(*messageToParse), &usableData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	_, found := usableData["WagonCommand"]
	if !found {
		return nil, errors.New("Unvalid Json")
	}
	return usableData, nil
}

func determineCommand(command uint8) (error, RecieveCommand) {
	var parsedCommand = RecieveCommand(command)
	switch parsedCommand {
	case SetTemp:
		return nil, parsedCommand
	case SetLight:
		return nil, parsedCommand
	default:
		return WrongGatewayCommand, 0
	}
}

func RecieverCommandPacket(message map[string]uint8, messageBuffPoint *[]byte) *[]byte {
	(*messageBuffPoint)[0] = message["TargetWagon"]
	(*messageBuffPoint)[1] = message["WagonCommand"]
	return messageBuffPoint
}

func CreateSetTempMess(message map[string]uint8) []byte {
	var messageBuff = make([]byte, 3)
	RecieverCommandPacket(message, &messageBuff)
	messageBuff[2] = message["TargetTemperature"]
	return messageBuff
}
func CreateSetLightMess(message map[string]uint8) []byte {
	var messageBuff = make([]byte, 3)
	RecieverCommandPacket(message, &messageBuff)
	messageBuff[2] = message["TargetTemperature"]
	return messageBuff
}
