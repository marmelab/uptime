package config

import (
	"testing"
)

func TestPassingInexistentConfigurationFileShouldLogError(t *testing.T) {
	_, err := GetConfig("/tmp/404")
	if err.Error() != "Configuration file not found: /tmp/404" {
		t.Errorf("Non existing configuration file should return an error. Got: ", err.Error())
	}
}

func TestParsingInvalidCnofigurationFileShouldLogError(t *testing.T) {
	_, err := GetConfig("./Configurator.go")
	if err.Error() != "Unable to parse JSON configuration: invalid character 'p' looking for beginning of value" {
		t.Errorf("Non existing configuration file should return an error. Got: ", err.Error())
	}
}

func TestGettingConfigurationShouldReturnMapOfValues(t *testing.T) {
	config, err := GetConfig("../conf.json.dist")

	if err != nil {
		t.Errorf("Parsing an existing and valid configuration file should not throw error. Got: %s", err.Error())
	}

	if config["port"].(string) != "8383" {
		t.Errorf("Retrieved configuration should return a map of configuration parameters. Got: config.port = ", config["port"])
	}
}
