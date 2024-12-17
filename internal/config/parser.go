package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ParseFileByPath(filePath string) (error, Config) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err, Config{}
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err, Config{}
	}

	return nil, config
}
