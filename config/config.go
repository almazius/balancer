package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	_config = Config{}
)

type Config struct {
	ApiServer   Server `json:"api-server"`
	ProxyServer Server `json:"proxy-server"`

	InitRoutes map[string][]string `json:"proxy-routes"`
}

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func ParseConfig() (*Config, error) {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("json")
	err := v.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("failed read config: %w", err)
	}

	err = v.Unmarshal(&_config)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshal config: %w", err)
	}

	return &_config, nil
}

func C() *Config {
	return &_config
}
