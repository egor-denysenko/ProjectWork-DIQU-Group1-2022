package protocolParser

import (
	"bytes"
	"testing"
)

func TestCreateSetTempMess(t *testing.T) {
	receiverTestCases := []struct {
		name     string
		mockData map[string]uint8
		want     []byte
	}{
		{name: "Verify Correct Json Creation", mockData: map[string]uint8{"WagonCommand": 70, "TargetWagon": 22, "TargetTemperature": 11}, want: []byte{22, 70, 11}},
		{name: "Correct Error Management", mockData: map[string]uint8{"WagonCommand": 70, "TargetWagon": 22, "TargetTemperatureee": 11}, want: nil},
	}
	for _, testCase := range receiverTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := createSetTempMess(testCase.mockData)
			if bytes.Compare(got, testCase.want) != 0 {
				t.Errorf("Running Test %v: \n Expected %v want %v", testCase.name, got, testCase.want)
			}
		})
	}
}
func TestCreateSetLightMess(t *testing.T) {
	receiverTestCases := []struct {
		name     string
		mockData map[string]uint8
		want     []byte
	}{
		{name: "Verify Correct Json Creation", mockData: map[string]uint8{"WagonCommand": 70, "TargetWagon": 22, "TargetLight": 11}, want: []byte{22, 70, 11}},
		{name: "Correct Error Management", mockData: map[string]uint8{"WagonCommand": 70, "TargetWagon": 22, "TargetLighteee": 11}, want: nil},
	}
	for _, testCase := range receiverTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := createSetLightMess(testCase.mockData)
			if bytes.Compare(got, testCase.want) != 0 {
				t.Errorf("Running Test %v: \n Expected %v want %v", testCase.name, got, testCase.want)
			}
		})
	}
}
func TestReceiverCommandPacket(t *testing.T) {
	receiverTestCases := []struct {
		name     string
		mockData map[string]uint8
		want     error
	}{
		{name: "Verify Correct Json Creation", mockData: map[string]uint8{"WagonCommand": 70, "TargetWagon": 22}, want: nil},
		{name: "Correct Error Management", mockData: map[string]uint8{"WagonCommand": 70, "TargetWagonnn": 22}, want: KeyDoesNotExistInMap},
	}
	for _, testCase := range receiverTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := receiverCommandPacket(testCase.mockData, &[]byte{1, 2})
			if got != testCase.want {
				t.Errorf("Running Test %v: \n Got %v Expected %v", testCase.name, got, testCase.want)
			}
		})
	}
}
func TestCheckIfKeyExists(t *testing.T) {
	receiverTestCases := []struct {
		name     string
		mockData map[string]uint8
		mockKey  string
		want     error
	}{
		{name: "Verify Correct Key Check", mockData: map[string]uint8{"WagonCommand": 70}, mockKey: "WagonCommand", want: nil},
		{name: "Correct Error Management", mockData: map[string]uint8{"WagonCommandd": 70}, mockKey: "WagonCommand", want: KeyDoesNotExistInMap},
	}
	for _, testCase := range receiverTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := checkIfKeyExists(testCase.mockData, testCase.mockKey)
			if got != testCase.want {
				t.Errorf("Running Test %v: \n Got %v Expected %v", testCase.name, got, testCase.want)
			}
		})
	}
}
