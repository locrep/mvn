package maven

import (
	"github.com/gin-gonic/gin"
	"github.com/locrep/go/config"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"os"
	"strings"
)

type handler interface {
	Handle(ctx *gin.Context)
}

type simpleHandler struct {
	envConf config.Environment
}

func NewSimpleHandler(envConf config.Environment) handler {
	return simpleHandler{envConf: envConf}
}

func (h simpleHandler) Handle(ctx *gin.Context) {
	for _, repo := range config.MavenOriginRepos {
		filePath := config.MavenRepo + ctx.Request.URL.String()

		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			ctx.File(filePath)
			return
		}

		response, body, errs := gorequest.New().Get(repo + ctx.Request.URL.String()).EndBytes()
		if len(errs) > 0 {
			ctx.JSON(http.StatusInternalServerError, DependencyFetchError(errs[0]))
			return
		} else if response.StatusCode != http.StatusOK {
			ctx.JSON(response.StatusCode, body)
			return
		}


		var (
			file *os.File
			err  error
		)

		paths := strings.Split(filePath, "/")
		fileName := paths[len(paths)-1]
		folder := filePath[0 : len(filePath)-len(fileName)]

		if err := os.MkdirAll(folder, 0777); err != nil {
			ctx.JSON(http.StatusInternalServerError, FileCreateError(err))
		}

		if file, err = os.Create(filePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, FileCreateError(err))
		}
		defer file.Close()

		if _, err := file.Write(body); err != nil {
			ctx.JSON(http.StatusInternalServerError, FileWriteError(err))
		}

		//todo: check sha and md5
		//todo: log

		ctx.File(filePath)
	}

}
