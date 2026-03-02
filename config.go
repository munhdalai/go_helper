package go_helper

import (
	"github.com/spf13/viper"
)

// LoadConfig нь тохиргооны файлаас уншиж, заасан struct руу unmarshal хийнэ.
//
// configName: файлын нэр (extension-гүй)
// configType: файлын төрөл (yaml, json, toml гэх мэт)
// path: файлын зам
func LoadConfig[T any](configName, configType, path string) (T, error) {
	var result T

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return result, err
	}

	if err := viper.Unmarshal(&result); err != nil {
		return result, err
	}

	return result, nil
}
