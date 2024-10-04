package configs

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type _Configuration struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	Database struct {
		Path      string `json:"path"`
		Migration struct {
			Path string `json:"path"`
		} `json:"migration"`
	} `json:"database"`
}

var configuration *_Configuration

func init() {
	if configuration != nil {
		return
	}

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(filepath.Join(basePath, "configs", "config.json"))
	if err != nil {
		panic(err)
	}

	configuration = new(_Configuration)
	if err := json.Unmarshal(content, configuration); err != nil {
		panic(err)
	}
}

func GetConfiguration() _Configuration {
	return *configuration
}
