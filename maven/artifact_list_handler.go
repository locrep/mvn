package maven

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/locrep/go/config"
)


//return artifact path like groupId/artifactId/version
func (a Artifact) String() string {
	return "/" + a.GroupID + "/" + a.ArtifactID + "/" + a.Version
}

type artifactHandler struct {
	envConf config.Environment
}

func NewArtifactHandler(envConf config.Environment) handler {
	return artifactHandler{envConf: envConf}
}

func (h artifactHandler) Handle(ctx *gin.Context) {
	fmt.Print("hello")
}
