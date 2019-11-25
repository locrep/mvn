package server

import (
	"github.com/gin-gonic/gin"
	"github.com/locrep/go/config"
	"github.com/locrep/go/maven"
)

type Handler interface {
	Handle(ctx *gin.Context)
}

func middleware(handler Handler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		handler.Handle(ctx)
		ctx.Next()
	}
}

func NewServer(envConf config.Environment) *gin.Engine {
	gin.SetMode(envConf.BuildMode)

	server := gin.New()
	mvnHandler := maven.NewHandler(envConf)
	server.Use(middleware(mvnHandler))

	return server
}
