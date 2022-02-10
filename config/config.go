package config

import (
	"github.com/alexflint/go-arg"
	log "github.com/sirupsen/logrus"
)

const RootPath = "/v1/palindrome"

type Args struct {
	LogLevel log.Level `arg:"env:LOG_LEVEL"`
	Port     int       `arg:"env:PORT"`
	Hostname string    `arg:"env:HOSTNAME,required"`
}

var Props Args

func LoadConfig() {
	arg.Parse(&Props)
}
