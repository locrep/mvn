package maven_test

import (
	. "github.com/locrep/mvn/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
)

var _ = Describe("when fetching existing artifact", func() {
	var (
		actualResp *http.Response
		err        error
	)

	BeforeAll(func() {
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
})
