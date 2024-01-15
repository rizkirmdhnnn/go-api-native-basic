package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	PORT            string
	DB_HOST         string
	DB_USER         string
	DB_PASSWORD     string
	DB_NAME         string
	DB_PORT         string
	PENALTYPERDAY   int32
	MAXLOANDURATION int32
	PENALTYBROKEN   int32
	PINALTYLOST     int32
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}

	log.Println("Load server successfully")
}
