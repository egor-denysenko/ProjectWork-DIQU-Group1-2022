package protocolParser

import (
	"testing"
)

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
