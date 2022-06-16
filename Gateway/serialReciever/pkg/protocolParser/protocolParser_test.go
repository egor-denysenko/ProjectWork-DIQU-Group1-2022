package protocolParser

import (
	"testing"
)

/*func TestParseSerialData(t *testing.T) {
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
}*/

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
			got := DetermineReciever(testCase.mockData)
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
			got, _ := DetermineCommand(testCase.mockData)
			if got != testCase.want {
				t.Errorf("Running Test %v: \n Expected %v want %v", testCase.name, got, testCase.want)
			}
		})
	}

}
