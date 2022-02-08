package main

import (
	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"github.com/namrahov/klawpse/config"
	"github.com/namrahov/klawpse/handler"
	"github.com/namrahov/klawpse/repo"
	"strconv"

	log "github.com/sirupsen/logrus"
	"net/http"
)

var opts struct {
	Profile string `short:"p" long:"profile" default:"dev" description:"Application run profile"`
}

func main() {

	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	initLogger()
	initEnvVars()
	config.LoadConfig()
	applyLogLevel()

	log.Info("Application is starting with profile: ", opts.Profile)

	err = repo.MigrateDb()
	if err != nil {
		log.Fatal(err)
	}
	repo.InitDb()

	router := mux.NewRouter()

	handler.ApplicationHandler(router)
	port := strconv.Itoa(config.Props.Port)
	log.Info("Starting server at port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
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

func applyLogLevel() {
	log.SetLevel(config.Props.LogLevel)
}
