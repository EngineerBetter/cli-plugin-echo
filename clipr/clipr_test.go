package clipr_test

import (
	. "github.com/EngineerBetter/cli-plugin-echo/clipr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bitly/go-simplejson"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("CLIPR", func() {
	Describe("/", func() {
		It("returns a listing containing the echo plugin", func() {
			server := httptest.NewServer(nil)

			indexHandler := IndexHandler{}
			indexHandler.Addr = server.URL
			server.Config.Handler = indexHandler
			defer server.Close()

			resp, err := http.Get(server.URL)
			Ω(err).ShouldNot(HaveOccurred())

			json, err := simplejson.NewFromReader(resp.Body)
			Ω(err).ShouldNot(HaveOccurred())

			echoNode := json.Get("plugins").GetIndex(0)
			Ω(echoNode.Get("name").MustString()).Should(Equal("echo"))
			bins, err := echoNode.Get("binaries").Array()
			Ω(err).ShouldNot(HaveOccurred())
			Ω(bins).Should(ContainElement(SatisfyAll(
				HaveKeyWithValue("platform", "osx"),
				HaveKeyWithValue("url", server.URL+"/bin/osx/echo"),
			)))
			Ω(bins).Should(ContainElement(SatisfyAll(
				HaveKeyWithValue("platform", "win64"),
				HaveKeyWithValue("url", server.URL+"/bin/windows64/echo.exe"),
			)))
		})
	})
})
