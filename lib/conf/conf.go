package conf

import (
	"encoding/json"
	"os"

	"github.com/farhaanbukhsh/file-indexer/lib/logger"
)

// Configuration loads the config file
type Configuration struct {
	RootDirectory string `json:"RootDirectory"`
	IndexFilename string `json:"IndexFilename"`
	Port          string `json:"Port"`
	LogFile       string `json:"LogFile"`
	LogMan        *logger.Logger
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
	config.LogMan = logger.NewLogger(config.LogFile, verbosity)
	return config
}
