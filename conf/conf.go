package conf

import (
	"encoding/json"
	"os"

	"github.com/farhaanbukhsh/file-indexer/logger"
)

// Configuration loads the config file
type Configuration struct {
	RootDirectory string
	IndexFilename string
	Port          string
}

// NewConfig is a function for creating a new configuration
func NewConfig(filename string, lg *logger.Logger) Configuration {
	config := Configuration{}
	name := "config.json"
	if filename != "" {
		name = filename
	}
	file, err := os.Open(name)
	defer file.Close()
	lg.Must(err, "Configuration File opened succesfully")
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	lg.Must(err, "JSON data successfully read")
	return config
}
