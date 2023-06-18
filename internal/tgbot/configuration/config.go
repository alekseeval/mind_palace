package configuration

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	System struct {
		MpAppSettings MindPalaceAppSettings `mapstructure:"mp_app"`
		Telegram      struct {
			AccessToken string `mapstructure:"access_token"`
		} `mapstructure:"telegram"`
	} `mapstructure:"system"`
	Logger struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"logger"`
}

type MindPalaceAppSettings struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func ReadConfig(path string) (config *Config, err error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
