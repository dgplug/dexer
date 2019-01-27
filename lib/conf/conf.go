package conf

import (
	"encoding/json"
	"os"

	"github.com/dgplug/dexer/lib/logger"
)

// Configuration loads the config file
type Configuration struct {
	RootDirectory string `json:"RootDirectory"`
	IndexFilename string `json:"IndexFilename"`
	Port          string `json:"Port"`
	LogFile       string `json:"LogFile"`
	cLogger       *logger.Logger
}

func (c *Configuration) Must(e error, logstring string) {
	c.cLogger.Must(e, logstring)
}

func (c *Configuration) GetLogger() *logger.Logger {
	return c.cLogger
}

// NewConfig is a function for creating a new configuration
func NewConfig(filename string, verbosity bool) Configuration {
	config := Configuration{}
	name := "config.json"
	if filename != "" {
		name = filename
	}
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	decoder.Decode(&config)
	config.cLogger = logger.NewLogger(config.LogFile, verbosity)
	return config
}
