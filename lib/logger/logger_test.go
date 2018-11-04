package logger

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLogWriter(t *testing.T) {
	filename := "testf"
	log := NewLogger(filename, false)
	log.Must(nil, "abc")

	input := formatLog("abc\n")
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		t.Errorf("log file with name testf not created")
	}

	testFormat(t, input, string(content))
	os.Remove(filename)
}

func testFormat(t *testing.T, input, output string) {

	if output[0] != input[0] {
		t.Errorf("output is in the wrong format, expected %b , got %b", input[0], output[0])
	}

	if output[26] != input[26] {
		t.Errorf("output is in the wrong format, expected %b , got %b", input[26], output[26])
	}

	if output[28:] != input[28:] {
		t.Errorf("output is in the wrong format, expected %s , got %s", input[28:], output[28:])
	}
}
