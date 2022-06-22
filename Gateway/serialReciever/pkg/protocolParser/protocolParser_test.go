package protocolParser

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestValidateSerialData(t *testing.T) {
	t.Run("test halal", func(t *testing.T) {
		mockDataChannel := make(chan []byte)
		mockRecievingChannel := make(chan []byte)
		mockData := []byte{31, 254, 69}
		desiredStruct := FormattedData{
			LocomotiveID: 188,
			VagonID:      31,
			Temperature:  30,
			Humidity:     80,
			VagonAllarms: VagonAllarms{
				DoorIO:         false,
				DoorB:          false,
				DoorC:          false,
				TemperatoreMin: false,
				TemperatureMax: false,
				Lights:         false,
				Humidity:       false,
			},
			VagonDoors: VagonDoors{
				Door1:       false,
				Door2:       false,
				Door3:       false,
				Door4:       false,
				DoorBath:    false,
				DoorConduct: false,
			},
			VagonLights: VagonLights{
				LightMode:   false,
				LightStatus: false,
			},
		}
		go func() {
			mockRecievingChannel <- mockData
			ValidateSerialData(mockRecievingChannel, mockDataChannel)
		}()
		got := <-mockRecievingChannel
		t.Log(string(got))
		var jsonToTest FormattedData
		err := json.Unmarshal(got, &jsonToTest)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		t.Log("json to test")
		t.Log(jsonToTest)
		log.Println("now getting data from channels")
		fmt.Println(desiredStruct)
		if !reflect.DeepEqual(desiredStruct, jsonToTest) {
			t.Errorf("error on serial validation")
		}
	})
}

func TestDetermineReciever(t *testing.T) {
	recieverTestCases := []struct {
		name     string
		mockData byte
		want     error
	}{
		{name: "Reciever Is Also The Id Of The Gateway", mockData: 254, want: nil},
		{name: "Reciever Is Not The Id Of The Gateway", mockData: 33, want: MessageNotForTheGateway},
	}
	for _, testCase := range recieverTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := determineReciever(testCase.mockData)
			if got != testCase.want {
				t.Errorf("Running Test %v: \n Expected %v want %v", testCase.name, got, testCase.want)
			}
		})
	}
}

func TestDetermineCommand(t *testing.T) {
	recieverTestCases := []struct {
		name     string
		mockData byte
		want     error
	}{
		{name: "Verify Correct Parsing Into Enum", mockData: 69, want: nil},
		{name: "Return Error Because The Value Is Not In The Enum", mockData: 33, want: WrongGatewayCommand},
	}
	for _, testCase := range recieverTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, _ := determineCommand(testCase.mockData)
			if got != testCase.want {
				t.Errorf("Running Test %v: \n Expected %v want %v", testCase.name, got, testCase.want)
			}
		})
	}
}
