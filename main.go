package main

import (
	"github.com/locrep/mvn/config"
	"github.com/locrep/mvn/db"
	"github.com/locrep/mvn/error"
	"github.com/locrep/mvn/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	envConf := config.Env()
	log.SetFormatter(&log.JSONFormatter{})

	client, err := db.NewClient(envConf.RedisUrl, envConf.RedisDB)
	if err != nil {
		log.WithFields(error.DefineError("Redis", 1, "Couldn't run server")(err)).Error(err.Error())
	}

	err = server.NewServer(envConf, client).Run(":" + envConf.Port)
	if err != nil {
		log.WithFields(error.DefineError("ServerError", 1, "Couldn't run server")(err)).Error(err.Error())
	}
}
