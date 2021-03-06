package config

import (
	"os"
	"time"
)

var MavenOriginRepos = []string{
	"https://repo.maven.apache.org/maven2",
}

var MavenRepo = "./repo"

type Environment struct {
	Port      string
	BuildMode string
	RedisUrl  string
	RedisDB   int
}

var Conf = struct {
	DBConnectionTimeout time.Duration
	DBReadTimeout       time.Duration
}{
	10, 10,
}

func Env() Environment {
	return Environment{
		Port:      os.Getenv("PORT"),
		BuildMode: os.Getenv("BUILD_MODE"),
		RedisUrl:  os.Getenv("REDIS_URL"),
		RedisDB:   0,
	}
}
