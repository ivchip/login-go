package utils

import (
	"errors"
	"github.com/spf13/viper"
)

func GetValueEnvironment(key string) (string, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		return "", errors.New("not found")
	}
	return value, nil
}