package api

import (
	"os"

	"gopkg.in/yaml.v2"
)

func ReadConfig(stPtr interface{}, path string) error {
	var file *os.File
	var err error
	file, err = os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(stPtr)
	if err != nil {
		return err
	}
	return nil
}
