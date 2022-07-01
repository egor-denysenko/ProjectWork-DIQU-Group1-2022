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
