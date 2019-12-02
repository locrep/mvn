package maven

import (
	"github.com/gin-gonic/gin"
	"github.com/locrep/mvn/config"
	"github.com/locrep/mvn/db"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"os"
	"strings"
)

type handler interface {
	Handle(ctx *gin.Context)
}

type simpleHandler struct {
	envConf  config.Environment
	dbClient *db.Client
}

func NewSimpleHandler(envConf config.Environment, dbClient *db.Client) handler {
	return simpleHandler{envConf: envConf, dbClient: dbClient}
}

func (h simpleHandler) Handle(ctx *gin.Context) {
	for _, repo := range config.MavenOriginRepos {
		filePath := config.MavenRepo + ctx.Request.URL.String()
		fileFolder := filePath[0:strings.LastIndex(filePath, "/")]
		redisKey := ctx.Request.URL.String()[0:strings.LastIndex(ctx.Request.URL.String(), "/")]
		redisValue := ctx.Request.URL.String()[strings.LastIndex(ctx.Request.URL.String(), "/"):]

		var (
			err                error
		)

		//if file exists then serve
		if _, err = os.Stat(filePath); !os.IsNotExist(err) {
			ctx.File(filePath)

			//is file in redis, if not, add it to redis
			go h.dbClient.Add(redisKey, redisValue)
		}

		response, body, errs := gorequest.New().Get(repo + ctx.Request.URL.String()).EndBytes()
		if len(errs) > 0 {
			ctx.JSON(http.StatusInternalServerError, DependencyFetchError(errs[0]))
			return
		} else if response.StatusCode != http.StatusOK {
			ctx.JSON(response.StatusCode, body)
			return
		}

		if err = os.MkdirAll(fileFolder, 0777); err != nil {
			ctx.JSON(http.StatusInternalServerError, FileCreateError(err))
		}

		var file *os.File
		if file, err = os.Create(filePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, FileCreateError(err))
		}
		defer file.Close()

		if _, err = file.Write(body); err != nil {
			ctx.JSON(http.StatusInternalServerError, FileWriteError(err))
		}

		//todo: check sha and md5
		//todo: log

		ctx.File(filePath)

		//is file in redis, if not, add it to redis
		go h.dbClient.Add(redisKey, redisValue)
	}

}
