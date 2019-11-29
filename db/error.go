package db

import . "github.com/locrep/mvn/error"

const ErrorPrefix = "REDIS"

var (
	CouldntConnectRedis     = DefineError(ErrorPrefix, 1, "Couldn't connect redis server")
	CouldntSetKeyValueData  = DefineError(ErrorPrefix, 2, "Couldn't set artifact key value data")
	CouldntReadKeyValueData = DefineError(ErrorPrefix, 3, "Couldn't read artifact key value data")
)
