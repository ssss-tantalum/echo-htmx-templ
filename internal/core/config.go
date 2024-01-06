package core

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Debug bool `mapstructure:"DEBUG"`

	SeverPort    int    `mapstructure:"SERVER_PORT"`
	ServerHost   string `mapstructure:"SERVER_HOST"`
	ReadTimeout  int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout int    `mapstructure:"WRITE_TIMEOUT"`

	DbDsn string `mapstructure:"DB_DSN"`
}

func LoadConfig() (*AppConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := new(AppConfig)
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}
