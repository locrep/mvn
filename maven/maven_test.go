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
	"os"
)

var _ = Describe("when artifact exists", func() {
	var (
		testServer      *httptest.Server
		actualResp      *http.Response
		err             error
		expectedContent = uuid.New().String()
	)
	const dummyArtifact = "/dummy-artifact.zip"

	BeforeAll(func() {
		if _, err = os.Stat(maven.Repo); os.IsNotExist(err) {
			err = os.Mkdir(maven.Repo, 0700)
			Expect(err).Should(BeNil())
		}

		_, err = os.Create(maven.Repo + dummyArtifact)
		Expect(err).Should(BeNil())

		err = ioutil.WriteFile(maven.Repo+dummyArtifact, []byte(expectedContent), 0644)
		Expect(err).Should(BeNil())

		envConf := config.Env()
		testServer = httptest.NewServer(server.NewServer(envConf))

		actualResp, err = testServer.Client().Get(testServer.URL + dummyArtifact)
		Expect(err).Should(BeNil())
	})

	It("the server should return 200 status ok", func() {
		Expect(actualResp.StatusCode).Should(Equal(http.StatusOK))
	})

	It("should return expected file", func() {
		actualContent, err := ioutil.ReadAll(actualResp.Body)
		Expect(err).Should(BeNil())

		Expect(actualContent).Should(Equal([]byte(expectedContent)))
	})

	AfterAll(func() {
		testServer.Close()

		err := os.RemoveAll(dummyArtifact)
		Expect(err).Should(BeNil())

		err = os.Remove(maven.Repo + dummyArtifact)
		Expect(err).Should(BeNil())
	})
})
