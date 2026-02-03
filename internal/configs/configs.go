package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MongoURI string `envconfig:"MONGO_URI"`
	MongoDB  string `envconfig:"MONGO_DB"`
	MongoCol string `envconfig:"MONGO_COL"`
}

var Env Config

func StartConfigs() error {
	if err := godotenv.Load("./.env"); err != nil {
		return err
	}

	if err := envconfig.Process("", &Env); err != nil {
		return err
	}

	return nil
}
