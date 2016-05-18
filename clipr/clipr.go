package clipr

import (
	"fmt"
	"net/http"
	"os"
)

type IndexHandler struct {
	Addr string
}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, responseBody, h.Addr, h.Addr)
}

func ServeOSX(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../bin/osx/echo")
}

func Configure(server *http.Server, addr string) {
	mux := http.NewServeMux()
	mux.Handle("/list", IndexHandler{Addr: addr})
	mux.HandleFunc("/bin/osx/echo", ServeOSX)
	server.Handler = mux
}

func main() {
	server := http.Server{}
	fmt.Println("Starting")
	err := server.ListenAndServe()
	Configure(&server, server.Addr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var responseBody = `{"plugins": [
  {
    "name":"echo",
    "description":"echo repeats input back to the terminal",
    "version":"0.1.4",
    "date":"0001-01-01T00:00:00Z",
    "company":"",
    "author":"",
    "contact":"feedback@email.com",
    "homepage":"https://github.com/johndoe/plugin-repo",
    "binaries": [
      {
        "platform":"osx",
        "url":"%s/bin/osx/echo",
        "checksum":"86aed94e9efd8bdda669c96fc36c979d9acbea5c"
      },
      {
        "platform":"win64",
        "url":"%s/bin/windows64/echo.exe",
        "checksum":"3062d690bc2991b93c29b823771c19257a7f42f5"
      }
    ]
  }
]}`
