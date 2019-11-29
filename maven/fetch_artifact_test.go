package maven_test

import (
	"github.com/google/uuid"
	"github.com/locrep/go/config"
	"github.com/locrep/go/maven"
	"github.com/locrep/go/server"
	. "github.com/locrep/go/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("when fetching existing artifact", func() {
	var (
		testServer      *httptest.Server
		actualResp      *http.Response
		err             error
		expectedContent = uuid.New().String()
		artifacts maven.Artifacts
	)

	BeforeAll(func() {
		artifacts = createDummyRepos([]byte(expectedContent))

		//run server
		envConf := config.Env()
		config.MavenRepo = dummyRepo
		testServer = httptest.NewServer(server.NewServer(envConf))

		artifactPath := artifacts[0].String()
		//fetch artifact
		actualResp, err = testServer.Client().Get(testServer.URL + artifactPath + dummyArtifact)
		Expect(err).Should(BeNil())
	})

	It("the server should return 200 status ok", func() {
		Expect(actualResp.StatusCode).Should(Equal(http.StatusOK))
	})

	It("the server should return expected file", func() {
		actualContent, err := ioutil.ReadAll(actualResp.Body)
		Expect(err).Should(BeNil())

		Expect(actualContent).Should(Equal([]byte(expectedContent)))
	})

	AfterAll(func() {
		testServer.Close()
		removeDummyArtifacts(artifacts)
	})
})
