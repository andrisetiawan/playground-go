package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

const CONFIG_FILE = "config.json"

type EnvironmentConfig struct {
	Development Config `json:"development"`
	Staging     Config `json:"staging"`
	Production  Config `json:"production"`
}

type Config struct {
	Port     string `json:port`
	Database string `json:database`
}

func (c *Config) LoadConfig(environment string) error {
	var configuration EnvironmentConfig

	file, _ := loadConfigFile()
	err := json.Unmarshal(file, &configuration)
	if err != nil {
		return err
	}

	switch strings.ToLower(environment) {
	case "development":
		*c = configuration.Development
	case "staging":
		*c = configuration.Staging
	case "production":
		*c = configuration.Production
	default:
		err = errors.New(fmt.Sprintf("Error: Environment '%s' is not found.", environment))
	}

	return err
}

func loadConfigFile() ([]byte, error) {
	return loadFile(CONFIG_FILE)
}

func loadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}
