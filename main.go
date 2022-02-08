package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var opts struct {
	Profile string `short:"p" long:"profile" default:"dev" description:"Application run profile"`
}

func main() {

	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	initLogger()
	initEnvVars()

}

func initLogger() {
	log.SetLevel(log.InfoLevel)
	if opts.Profile == "dev" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func initEnvVars() {
	if godotenv.Load("profiles/dev.env") != nil {
		log.Fatal("Error in loading environment variables from: profiles/dev.env")
	} else {
		log.Info("Environment variables loaded from: profiles/dev.env")
	}

	if opts.Profile != "dev" {
		profileFileName := "profiles/" + opts.Profile + ".env"
		if godotenv.Overload(profileFileName) != nil {
			log.Fatal("Error in loading environment variables from: ", profileFileName)
		} else {
			log.Info("Environment variables overloaded from: ", profileFileName)
		}
	}
}
