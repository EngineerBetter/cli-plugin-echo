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
			server := httptest.NewServer(http.HandlerFunc(ServeIndex))
			defer server.Close()

			resp, err := http.Get(server.URL)
			Ω(err).ShouldNot(HaveOccurred())

			json, err := simplejson.NewFromReader(resp.Body)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(json.Get("plugins").GetIndex(0).Get("name").MustString()).Should(Equal("echo"))
		})
	})
})
