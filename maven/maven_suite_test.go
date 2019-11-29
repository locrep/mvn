package maven_test

import (
	"github.com/google/uuid"
	"github.com/locrep/mvn/config"
	"github.com/locrep/mvn/db"
	"github.com/locrep/mvn/maven"
	"github.com/locrep/mvn/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	dummyRepo     = "./dummy-repo"
	dummyArtifact = "/dummy-artifact.zip"

	testServer      *httptest.Server
	actualResp      *http.Response
	err             error
	expectedContent = uuid.New().String()
	artifacts       maven.Artifacts
	redis           *db.Client
)

const testDB = 1

func TestMaven(t *testing.T) {
	config.MavenRepo = dummyRepo

	RegisterFailHandler(Fail)
	RunSpecs(t, "Maven Suite")
}

var _ = BeforeSuite(func() {
	artifacts = createDummyRepos([]byte(expectedContent))
	redis = createDummyRepoKeyValue(artifacts)

	//run server
	envConf := config.Env()
	config.MavenRepo = dummyRepo
	testServer = httptest.NewServer(server.NewServer(envConf, redis))
})

var _ = AfterSuite(func() {
	testServer.Close()
	removeDummyArtifacts(artifacts)
	removeAllKeyValues(redis)
})

func createDummyRepoKeyValue(artifacts maven.Artifacts) *db.Client {
	client, err := db.NewClient(config.Env().RedisUrl, testDB)
	Expect(err).Should(BeNil())

	for _, artifact := range artifacts {
		err = client.Add(artifact.String(), dummyArtifact)
		Expect(err).Should(BeNil())
	}

	return client
}

func removeAllKeyValues(client *db.Client) {
	err := client.Redis.FlushDB().Err()
	Expect(err).Should(BeNil())
}

func createDummyRepos(artifact []byte) maven.Artifacts {
	var err error
	dummyRepos := maven.Artifacts{
		maven.Artifact{
			GroupID:    "ant",
			ArtifactID: "ant",
			Version:    "1.7.0",
		},
		maven.Artifact{
			GroupID:    "asm",
			ArtifactID: "asm-parent",
			Version:    "3.3.1",
		},
	}

	for _, repo := range dummyRepos {
		repoPath := dummyRepo + repo.String()
		if _, err = os.Stat(repoPath); os.IsNotExist(err) {
			err = os.MkdirAll(repoPath, 0700)
			Expect(err).Should(BeNil())
		}

		_, err = os.Create(repoPath + dummyArtifact)
		Expect(err).Should(BeNil())

		err = ioutil.WriteFile(repoPath+dummyArtifact, artifact, 0644)
		Expect(err).Should(BeNil())
	}

	return dummyRepos
}

func removeDummyArtifacts(artifacts maven.Artifacts) {
	for _, artifact := range artifacts {
		err := os.RemoveAll(config.MavenRepo + artifact.String())
		Expect(err).Should(BeNil())
	}

	err := os.RemoveAll(config.MavenRepo)
	Expect(err).Should(BeNil())
}
