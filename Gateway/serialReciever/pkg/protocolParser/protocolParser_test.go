package protocolParser

import (
	"context"
	"testing"
)

func TestValidateSerialData(t *testing.T) {
	validationTestCases := []struct {
		name                 string
		mockDataChannel      chan []byte
		mockRecievingChannel chan []byte
		mockData             []byte
		want                 FormattedData
	}{
		{
			name:                 "Test Value Reciever Is Uncorrect",
			mockDataChannel:      make(chan []byte),
			mockRecievingChannel: make(chan []byte),
			mockData:             []byte{31, 254, 69},
			want: FormattedData{
				locomotiveID: 0,
				vagonID:      31,
				temperature:  30,
				humidity:     80,
				vagonAllarms: vagonAllarms{},
				vagonDoors:   vagonDoors{},
				vagonLights:  vagonLights{},
			},
		},
	}
	for _, testCases := range validationTestCases {
		t.Run(testCases.name, func(t *testing.T) {
			ctx := context.Background()
			ctx, cancelCtx := context.WithCancel(ctx)
			defer cancelCtx()
			go ValidateSerialData(ctx, testCases.mockRecievingChannel, testCases.mockDataChannel)
		})
	}
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
