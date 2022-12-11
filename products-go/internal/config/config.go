package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	configPath = "configs/application.yml"

	parameterEnvPrefix = "${"
	parameterEnvSuffix = "}"
)

type (
	Config struct {
		Application ApplicationConfig `json:"application" yaml:"application"`
		Datasource  DatasourceConfig  `json:"datasource" yaml:"datasource"`
		Server      ServerConfig      `json:"server" yaml:"server"`
	}

	ApplicationConfig struct {
		Name string
	}

	DatasourceConfig struct {
		Dsn string
	}

	ServerConfig struct {
		Port uint
	}
)

func Init() (*Config, error) {
	setUpConfigData(configPath)

	readingConfigError := viper.ReadInConfig()
	if readingConfigError != nil {
		return nil, readingConfigError
	}

	bindingEnvError := bindEnvToParameters()
	if bindingEnvError != nil {
		return nil, bindingEnvError
	}

	config := &Config{}

	unmarshallingError := viper.Unmarshal(config)
	if unmarshallingError != nil {
		return nil, unmarshallingError
	}

	return config, nil
}

func setUpConfigData(path string) {
	viper.SetConfigFile(path)
}

func bindEnvToParameters() error {
	for _, parameterKey := range viper.AllKeys() {
		parameterValue := viper.GetString(parameterKey)
		parameterEnv := isParameterEnv(parameterValue)

		if parameterEnv {
			envName := getEnvNameFromParameter(parameterValue)
			envValue, gettingEnvValueError := getEnvValue(envName)

			if gettingEnvValueError != nil {
				return gettingEnvValueError
			}

			viper.Set(parameterKey, envValue)
		}
	}

	return nil
}

func isParameterEnv(parameter string) bool {
	parameterHasEnvPrefix := strings.HasPrefix(parameter, parameterEnvPrefix)
	parameterHasEnvSuffix := strings.HasSuffix(parameter, parameterEnvSuffix)

	return parameterHasEnvPrefix && parameterHasEnvSuffix
}

func getEnvNameFromParameter(parameter string) string {
	parameterWithoutEnvPrefix := strings.TrimPrefix(parameter, parameterEnvPrefix)
	return strings.TrimSuffix(parameterWithoutEnvPrefix, parameterEnvSuffix)
}

func getEnvValue(envName string) (string, error) {
	envValue := os.Getenv(envName)
	if len(envValue) == 0 {
		errorText := fmt.Sprintf("%s is not declared", envName)
		return "", errors.New(errorText)
	}

	return envValue, nil
}
