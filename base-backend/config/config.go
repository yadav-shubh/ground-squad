package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	Server struct {
		Host string
	}
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		DBName   string
		SSLMode  string
	}
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Failed to load config", err)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		log.Println("Failed to unmarshal config", err)
	}

	cfg = &c
}

func Get() *Config {
	once.Do(loadConfig)
	return cfg
}
