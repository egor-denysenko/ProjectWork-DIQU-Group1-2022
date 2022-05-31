package protocolParser

import (
	"bytes"
	"testing"
)

func TestParseSerialData(t *testing.T) {
	testBytes := []byte{0, 1, 2, 4}
	got := ParseSerialData(testBytes)

	wanted := bytes.Compare(testBytes, got)
	if wanted != 0 {
		t.Errorf("Wanted %q, but parsed %q", testBytes, got)
	}
}
