package config

import (
    "fmt"

    "github.com/spf13/viper"
)

func LoadConfig[T any](configPath, configName, configType string) (*T, error) {
    var config T

    viper.AddConfigPath(configPath)
    viper.SetConfigName(configName)
    viper.SetConfigType(configType)

    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("failed to read the config file")
    }

    if err := viper.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("failed to unmarshal the config file")
    }

    return &config, nil
}