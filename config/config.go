package config

import (
	"os"
)

var MavenRepos = []string{
	"https://repo.maven.apache.org/maven2",
}

type Environment struct {
	Port      string
	BuildMode string
}

func Env() Environment {
	return Environment{
		Port:      os.Getenv("PORT"),
		BuildMode: os.Getenv("BUILD_MODE"),
	}
}
