package config

import (
	"os"
)

var MavenOriginRepos = []string{
	"https://repo.maven.apache.org/maven2",
}

var MavenRepo = "./repo"

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
