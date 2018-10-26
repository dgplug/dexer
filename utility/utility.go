package utility

import (
	"fmt"
	"io/ioutil"
	"os"
)

// GetContent is a function for retrieving data from file
func GetContent(name string) (string, error) {
	data, err := ioutil.ReadFile(name)
	return string(data), err
}

// DeleteExistingIndex checks if the index exist if it does, then flushes it off
func DeleteExistingIndex(name string) error {
	_, err := os.Stat(name)
	if !os.IsNotExist(err) {
		if err := os.RemoveAll(name); err != nil {
			return fmt.Errorf("Can't Delete file: %v", err)
		}
	}
	return nil
}
