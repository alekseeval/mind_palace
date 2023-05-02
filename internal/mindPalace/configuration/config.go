package configuration

import (
	"github.com/spf13/viper"
)

type Config struct {
	System struct {
		Http struct {
			Host string `mapstructure:"host"`
			Port int    `mapstructure:"port"`
		}
		Grpc struct {
			Port int `mapstructure:"port"`
		}
		DB struct {
			Host    string `mapstructure:"host"`
			Port    int    `mapstructure:"port"`
			Name    string `mapstructure:"name"`
			MaxConn int    `mapstructure:"max_conn"`
		}
	}
}

func ReadConfig(path string) (config *Config, err error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
