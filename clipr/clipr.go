package clipr

import (
	"fmt"
	"net/http"
	"os"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, responseBody)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(ServeIndex))
	server := http.Server{Handler: mux}
	fmt.Println("Starting")
	err := server.ListenAndServe()

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
        "url":"https://github.com/johndoe/plugin-repo/raw/master/bin/osx/echo",
        "checksum":"2a087d5cddcfb057fbda91e611c33f46"
      },
      {
        "platform":"win64",
        "url":"https://github.com/johndoe/plugin-repo/raw/master/bin/windows64/echo.exe",
        "checksum":"b4550d6594a3358563b9dcb81e40fd66"
      }
    ]
  }
]}`
