package maven

import (
	"github.com/gin-gonic/gin"
	"github.com/locrep/go/config"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"os"
	"strings"
)

const Repo = "./repo"

type handler struct {
	envConf config.Environment
}

func NewHandler(envConf config.Environment) handler {
	return handler{envConf: envConf}
}

func (h handler) Handle(ctx *gin.Context) {
	for _, repo := range config.MavenRepos {
		filePath := Repo + ctx.Request.URL.String()

		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			ctx.File(filePath)
			return
		}

		response, body, errs := gorequest.New().Get(repo + ctx.Request.URL.String()).EndBytes()
		if len(errs) > 0 || response.StatusCode != http.StatusOK {
			ctx.JSON(http.StatusNotFound, DependencyFetchError(errs[0]))
		}

		var (
			file *os.File
			err  error
		)

		paths := strings.Split(filePath, "/")
		fileName := paths[len(paths)-1]
		folder := filePath[0 : len(filePath)-len(fileName)]

		if err := os.MkdirAll(folder, 0777); err != nil {
			ctx.JSON(http.StatusNotFound, FileCreateError(err))
		}

		if file, err = os.Create(filePath); err != nil {
			ctx.JSON(http.StatusNotFound, FileCreateError(err))
		}
		defer file.Close()

		if _, err := file.Write(body); err != nil {
			ctx.JSON(http.StatusNotFound, FileWriteError(err))
		}
		//todo: check sha and md5
		//todo: log

		ctx.File(filePath)
	}

}
