package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
)

type MyHandler struct {
}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)

	// data, err := ioutil.ReadFile(string(path))
	f, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(f)
		
		var contentType string
		if strings.HasSuffix(path, ".css"){
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else if strings.HasSuffix(path, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 My Friend - " + http.StatusText(404)))
	}
}
func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8080", nil)
}