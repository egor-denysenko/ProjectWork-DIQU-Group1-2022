package protocolParser

import (
	"testing"
)

func TestParseSerialData(t *testing.T) {
	serialParserTests := []struct{
		name       string
		mockData []byte
		want       []byte
	}
	for _,testCase := range serialParserTests{
		t.Run(testCase.name,func(t *testing.T){
			got,_ := ParseSerialData(testCase.mockData)
		})
	}
}

func TestParseSerialDataErr(t *testing.T) {
	serialParserTests := []struct{
		name       string
		mockData []byte
		want       error
	}
	for _,testCase := range serialParserTests{
		t.Run(testCase.name,func(t *testing.T){
			_,err := ParseSerialData(testCase.mockData)
		})
	}

}