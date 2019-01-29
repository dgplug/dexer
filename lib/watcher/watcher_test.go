package watcher

import (
	"os"
	"testing"
	"time"

	"github.com/dgplug/dexer/lib/conf"
)

func TestWatcher(t *testing.T) {
	createConfig(t)
	config := conf.NewConfig("config.json", false)
	duration := 10 * time.Millisecond
	w := NewWatcher(config, duration)
	action := func(path string, status FileStatus) {
		t.Log(path, status)
	}

	go w.Start(action)

	doChanges(t)
	cleanUp()
}

func createConfig(t *testing.T) {
	data := `{
	"RootDirectory": "logs",
	"IndexFilename": "irclogs.bleve",
	"Port": ":8000",
	"LogFile": "logfile"
}`

	os.Mkdir("logs", os.FileMode(0777))
	file, err := os.Create("config.json")

	if err != nil {
		cleanUp()
		t.Fatalf("Unable to create the config.")
	}

	_, err = file.WriteString(data)

	if err != nil {
		cleanUp()
		t.Fatalf("Unable to write to the config.")
	}

	file.Close()
}

func doChanges(t *testing.T) {
	file, err := os.Create("logs/test")

	if err != nil {
		cleanUp()
		t.Fatalf("Unable to create a temoparary file.")
	}

	file.Close()

	file, err = os.OpenFile("logs/test", os.O_RDWR, 0644)

	if err != nil {
		cleanUp()
		t.Fatalf("Unable to create a temoparary file.")
	}

	_, err = file.WriteString("testing")

	if err != nil {
		cleanUp()
		t.Fatalf("Unable to write to the temporary file.")
	}

	file.Close()

	err = os.Remove("logs/test")
	if err != nil {
		cleanUp()
		t.Fatalf("Unable to remove the file.")
	}
}

func cleanUp() {
	os.Remove("config.json")
	os.Remove("logfile")
	os.RemoveAll("logs")
}
