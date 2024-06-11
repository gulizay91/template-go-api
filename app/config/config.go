package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	Server  ServerConfig  `mapstructure:"server"`
	Service ServiceConfig `mapstructure:"service"`
}

func (config *Config) Validate() error {
	err := validation.ValidateStruct(
		config,
		validation.Field(&config.Server),
		validation.Field(&config.Service),
	)
	return err
}

type ServerConfig struct {
	Port string
	Addr string
}

func (config ServerConfig) Validate() error {
	err := validation.ValidateStruct(
		&config,
		validation.Field(&config.Addr, validation.Required),
		validation.Field(&config.Port, is.Port),
	)
	return err
}

type ServiceConfig struct {
	LogLevel    string `mapstructure:"logLevel"`
	Name        string
	Environment string
}

func (config ServiceConfig) Validate() error {
	err := validation.ValidateStruct(
		&config,
		validation.Field(&config.LogLevel, validation.Required),
		validation.Field(&config.Environment, validation.Required),
	)
	return err
}
