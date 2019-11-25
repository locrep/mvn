package main

import (
	"github.com/locrep/go/config"
	"github.com/locrep/go/server"
)

func main() {
	envConf := config.Env()

	server.NewServer(envConf).Run(":" + envConf.Port)
}
