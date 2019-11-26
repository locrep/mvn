package maven_test

import (
	"github.com/locrep/go/config"
	"github.com/locrep/go/maven"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	dummyRepo     = "./dummy-repo"
	dummyArtifact = "/dummy-artifact.zip"
)

func TestMaven(t *testing.T) {
	config.MavenRepo = dummyRepo

	RegisterFailHandler(Fail)
	RunSpecs(t, "Maven Suite")
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

		err = ioutil.WriteFile(repoPath + dummyArtifact, artifact, 0644)
		Expect(err).Should(BeNil())
	}

	return dummyRepos
}

func removeDummyArtifacts(artifacts maven.Artifacts) {
	for _, artifact := range artifacts {
		err := os.RemoveAll(artifact.String())
		Expect(err).Should(BeNil())
	}
}
