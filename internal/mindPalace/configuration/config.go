package configuration

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	System struct {
		Http HttpConfig `mapstructure:"http"`
		Grpc struct {
			Port int `mapstructure:"port"`
		} `mapstructure:"grpc"`
		DB struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			DBName   string `mapstructure:"db_name"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			MaxConn  int    `mapstructure:"max_conn"`
			Timeout  int    `mapstructure:"timeout"`
		} `mapstructure:"db"`
	} `mapstructure:"system"`
	Logger struct {
		Level string `mapstructure:"level"`
	}
}

type HttpConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
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
