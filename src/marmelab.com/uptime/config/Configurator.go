package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func GetConfig(configPath string) (map[string]interface{}, error) {
	rawConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errors.New("Configuration file not found: " + configPath)
	}

	var config map[string]interface{}
	err = json.Unmarshal(rawConfig, &config)
	if err != nil {
		return nil, errors.New("Unable to parse JSON configuration: " + err.Error())
	}

	return config, nil
}
