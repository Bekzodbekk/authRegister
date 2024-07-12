package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type MongoConfig struct {
	Host string
	Port string
	Database string
	Collection string
}

type Config struct {
	Mongosh MongoConfig

	ServiceHost string
	ServicePort string
}

func Load(path string) (*Config, error) {
	
	err := godotenv.Load(path + "/.env")
	if err != nil {
		return nil, err
	}

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		Mongosh: MongoConfig{
			Host: conf.GetString("MONGOSH_HOST"),
			Port: conf.GetString("MONGOSH_PORT"),
			Database: conf.GetString("MONGOSH_DATABASE"),
			Collection: conf.GetString("MONGOSH_COLLECTION"),
		},

		ServiceHost: conf.GetString("SERVICE_HOST"),
		ServicePort: conf.GetString("SERVICE_PORT"),
	}

	if cfg.Mongosh.Database == "" || cfg.Mongosh.Collection == "" || cfg.ServiceHost == "" || cfg.ServicePort == "" {
		return nil, errors.New("FAILED while getting from .env")
	}

	return &cfg, nil
}
