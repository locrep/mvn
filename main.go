package main

import (
	"github.com/locrep/go/config"
	"github.com/locrep/go/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	envConf := config.Env()
	log.SetFormatter(&log.JSONFormatter{})

	server.NewServer(envConf).Run(":" + envConf.Port)
}
