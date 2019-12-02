package maven

import (
	"github.com/gin-gonic/gin"
	"github.com/locrep/mvn/config"
	"github.com/locrep/mvn/db"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

type artifactHandler struct {
	envConf  config.Environment
	dbClient *db.Client
}

func NewArtifactHandler(envConf config.Environment, dbClient *db.Client) handler {
	return artifactHandler{envConf: envConf, dbClient: dbClient}
}

func (h artifactHandler) Handle(ctx *gin.Context) {
	// Scan all keys
	var cursor uint64
	var err error

	artifactRepos := make(ArtifactRepos, 0)
	for {
		var keys []string
		if keys, cursor, err = h.dbClient.Redis.Scan(cursor, "", 50).Result(); err != nil {
			ctx.JSON(http.StatusInternalServerError, DependencyFetchError(err))
			return
		}

		if len(keys) <= 0 {
			ctx.JSON(http.StatusNotFound, "There is no artifact")
			return
		}

		for _, key := range keys {
			files, err := h.dbClient.Get(key)
			if err != nil {
				logger.WithFields(ThereIsNoArtifact(err)).Error(err.Error())
				continue
			}

			artifactRepos = append(artifactRepos, ArtifactRepo{
				Repo:  key,
				Files: files,
			})

		}

		if cursor == 0 {
			break
		}
	}

	ctx.JSON(http.StatusOK, artifactRepos)
}
