package config

import (
	"github.com/alexflint/go-arg"
	log "github.com/sirupsen/logrus"
)

const RootPath = "/v1/bracket"

type Args struct {
	DbHost   string    `arg:"env:DB_CARD_DELIVERY_HOST,required"`
	DbPort   string    `arg:"env:DB_CARD_DELIVERY_PORT,required"`
	DbName   string    `arg:"env:DB_CARD_DELIVERY_NAME,required"`
	DbUser   string    `arg:"env:DB_CARD_DELIVERY_USER,required"`
	DbPass   string    `arg:"env:DB_CARD_DELIVERY_PASS,required"`
	LogLevel log.Level `arg:"env:LOG_LEVEL"`
	Port     int       `arg:"env:PORT"`
	Hostname string    `arg:"env:HOSTNAME,required"`
}

var Props Args

func LoadConfig() {
	arg.Parse(&Props)
}
