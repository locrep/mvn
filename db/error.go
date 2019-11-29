package db

import "fmt"

const MongoErrorsPrefix = "MON"

var (
	CouldntConnectMongoServer       = defineError(1, "Could'nt connect mongo server")
	CouldntConnectMongoDB           = defineError(2, "Could'nt connect mongo db")
	CouldntReadArtifactsFromMongoDB = defineError(3, "Could'nt read artifacts from mongo db")
	CouldntDecodeArtifact           = defineError(4, "Could'nt decode artifact")
	GotErrorFromArtifactCursor      = defineError(5, "Got error from artifact cursor")
)

func defineError(index int, msg string) func(error) map[string]interface{} {
	return func(err error) map[string]interface{} {
		return map[string]interface{}{
			"code":    fmt.Sprintf("%s-%03d", MongoErrorsPrefix, index),
			"message": msg,
			"cause":   err.Error(),
		}
	}
}
