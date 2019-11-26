package maven_test

import (
	"encoding/json"
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

var _ = Describe("When try to get artifact list", func() {
	var (
		testServer        *httptest.Server
		actualResp        *http.Response
		err               error
		expectedArtifacts maven.Artifacts
	)

	BeforeAll(func() {
		expectedArtifacts = createDummyRepos([]byte(uuid.New().String()))

		envConf := config.Env()
		testServer = httptest.NewServer(server.NewServer(envConf))

		actualResp, err = testServer.Client().Get(testServer.URL + "/artifact")
		Expect(err).Should(BeNil())
	})

	It("the server should return 200 status ok", func() {
		Expect(actualResp.StatusCode).Should(Equal(http.StatusOK))
	})

	It("should return all artifacts", func() {
		actualContent, err := ioutil.ReadAll(actualResp.Body)
		Expect(err).Should(BeNil())

		var actualArtifacts maven.Artifacts
		err = json.Unmarshal(actualContent, &actualArtifacts)
		Expect(err).Should(BeNil())

		Expect(actualArtifacts).Should(Equal(expectedArtifacts))
	})

	AfterAll(func() {
		testServer.Close()
		removeDummyArtifacts(expectedArtifacts)
	})
})
