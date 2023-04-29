package config

import (
	"log"

	"github.com/spf13/viper"
)

var Port = ":8080"

type Config struct {
	// Your configuration fields here
}

func Load() (*Config, error) {
	viper.SetConfigName("config/config") // Name of config file (without extension)
	viper.SetConfigType("yaml")          // Type of config file
	viper.AddConfigPath(".")             // Search the root directory for the config file

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}
