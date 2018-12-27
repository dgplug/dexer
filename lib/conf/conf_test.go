package conf

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestConfReader(t *testing.T) {
	filedata := []byte(`{
	"RootDirectory": "logs",
	"IndexFilename": "irclogs.bleve",
	"Port": ":8000",
	"LogFile": "logfile"
}`)

	err := ioutil.WriteFile("config.json", filedata, 0777)
	if err != nil {
		t.Fatalf("Not able to create a config file for testing.")
	}

	expected := struct {
		root    string
		index   string
		port    string
		logfile string
	}{
		root:    "logs",
		index:   "irclogs.bleve",
		port:    ":8000",
		logfile: "logfile",
	}

	config := NewConfig("config.json", false)

	if config.RootDirectory != expected.root {
		t.Errorf("Root Directory not detected or has problems with it.")
	}

	if config.IndexFilename != expected.index {
		t.Errorf("Index File Name not detected or has problems with it.")
	}

	if config.Port != expected.port {
		t.Errorf("Port not detected or has problems with it.")
	}

	if config.LogFile != expected.logfile {
		t.Errorf("Log File not detected or has problems with it.")
	}

	err = os.Remove("config.json")

	if err != nil {
		t.Logf("Unable to remove config.json, please clean it up manually")
	}
}
