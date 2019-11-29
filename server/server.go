package server

import (
	"github.com/gin-gonic/gin"
	"github.com/locrep/mvn/config"
	"github.com/locrep/mvn/db"
	"github.com/locrep/mvn/maven"
	"strings"
)

type Handler interface {
	Handle(ctx *gin.Context)
}

func middleware(handler Handler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if !strings.Contains(ctx.Request.URL.Path, "v1") {
			handler.Handle(ctx)
			ctx.Next()
		}
	}
}

func NewServer(envConf config.Environment, client *db.Client) *gin.Engine {
	gin.SetMode(envConf.BuildMode)

	server := gin.New()
	server.Use(gin.Recovery())
	mvnHandler := maven.NewSimpleHandler(envConf, client)
	server.Use(middleware(mvnHandler))

	server.GET("/v1/artifact/", maven.NewArtifactHandler(envConf, client).Handle)
	return server
}
