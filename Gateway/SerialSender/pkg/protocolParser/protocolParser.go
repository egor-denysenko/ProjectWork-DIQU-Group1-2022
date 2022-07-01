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
var KeyDoesNotExistInMap = errors.New("the key doesn't exist in the provided map")
var UnvalidJson = errors.New("Unvalid Json")

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
		parsedDataChan <- createSetTempMess(data)
	case SetLight:
		parsedDataChan <- createSetLightMess(data)
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
	errTargetWagon := checkIfKeyExists(usableData, "TargetWagon")
	if errTargetWagon != nil {
		return nil, UnvalidJson
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

func receiverCommandPacket(message map[string]uint8, messageBuffPoint *[]byte) error {
	errWagonCommand := checkIfKeyExists(message, "WagonCommand")
	errTargetWagon := checkIfKeyExists(message, "TargetWagon")
	if errWagonCommand != nil || errTargetWagon != nil {
		return KeyDoesNotExistInMap
	}
	(*messageBuffPoint)[0] = message["TargetWagon"]
	(*messageBuffPoint)[1] = message["WagonCommand"]
	return nil
}

func createSetTempMess(message map[string]uint8) []byte {
	var messageBuff = make([]byte, 3)
	err := receiverCommandPacket(message, &messageBuff)
	if err != nil {
		return nil
	}
	errTargetTemperature := checkIfKeyExists(message, "TargetTemperature")
	if errTargetTemperature != nil {
		return nil
	}
	messageBuff[2] = message["TargetTemperature"]
	return messageBuff
}

func createSetLightMess(message map[string]uint8) []byte {
	var messageBuff = make([]byte, 3)
	err := receiverCommandPacket(message, &messageBuff)
	if err != nil {
		return nil
	}
	errTargetTemperature := checkIfKeyExists(message, "TargetLight")
	if errTargetTemperature != nil {
		return nil
	}
	messageBuff[2] = message["TargetLight"]
	return messageBuff
}

func checkIfKeyExists(message map[string]uint8, key string) error {
	if _, ok := message[key]; ok {
		return nil
	}
	return KeyDoesNotExistInMap
}
