package server

import (
	"github.com/gin-gonic/gin"
	"github.com/locrep/go/config"
	"github.com/locrep/go/maven"
	"strings"
)

type Handler interface {
	Handle(ctx *gin.Context)
}

func middleware(handler Handler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if !strings.Contains(ctx.Request.URL.Path,"v1"){
			handler.Handle(ctx)
			ctx.Next()
		}
	}
}

func NewServer(envConf config.Environment) *gin.Engine {
	gin.SetMode(envConf.BuildMode)

	server := gin.New()
	server.Use(gin.Recovery())
	mvnHandler := maven.NewSimpleHandler(envConf)
	server.Use(middleware(mvnHandler))

	server.GET("/v1/artifact/",maven.NewArtifactHandler(envConf).Handle)
	return server
}
