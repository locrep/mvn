package maven

import "github.com/locrep/mvn/error"

const ErrorPrefix = "MVN"

var (
	DependencyFetchError = error.DefineError(ErrorPrefix, 1, "Couldn't fetch dependency")
	FileCreateError      = error.DefineError(ErrorPrefix, 2, "Couldn't create file path")
	FileWriteError       = error.DefineError(ErrorPrefix, 3, "Couldn't write file")
	ThereIsNoArtifact    = error.DefineError(ErrorPrefix, 4, "There is no artifact")
)
