package maven_test

import (
	"encoding/json"
	"github.com/locrep/mvn/maven"
	. "github.com/locrep/mvn/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
)

var _ = Describe("When try to get artifact list", func() {
	var (
		actualResp        *http.Response
		err               error
		expectedArtifacts maven.ArtifactRepos
	)

	BeforeAll(func() {
		expectedArtifacts = make(maven.ArtifactRepos, 0)
		for _, artifact := range artifacts {
			artifactRepo := maven.ArtifactRepo{
				Repo:  artifact.String(),
				Files: []string{dummyArtifact},
			}

			expectedArtifacts = append(expectedArtifacts, artifactRepo)
		}
		actualResp, err = testServer.Client().Get(testServer.URL + "/v1/artifact")
		Expect(err).Should(BeNil())
	})

	It("the server should return 200 status ok", func() {
		Expect(actualResp.StatusCode).Should(Equal(http.StatusOK))
	})

	It("should return all artifacts", func() {
		actualContent, err := ioutil.ReadAll(actualResp.Body)
		Expect(err).Should(BeNil())

		var actualArtifactRepos maven.ArtifactRepos
		err = json.Unmarshal(actualContent, &actualArtifactRepos)
		Expect(err).Should(BeNil())

		Expect(len(actualArtifactRepos)).Should(Equal(len(expectedArtifacts)))
	})
})
