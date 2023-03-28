package main

import (
	"net/http"
	"text/template"
)

type Context struct {
    FirstName string
    Message string
	URL string
}

func main() {
    http.HandleFunc("/", myHandlerFunc)
    http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        context := Context{"Todd", "more Go, please", req.URL.Path}
        tmpl.Execute(w, context)
    }
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>First Template</title>
</head>
<body>
    {{ if eq .URL "/nobeer" }}
		<h1>We're out of beer. Sorry!</h1>
	{{ else }}
		<h1>Yes, grab another beer, {{ .FirstName }}</h1>
	{{ end }}

	<hr>
	<h2>Here is all the data:</h2>
	<p>{{.}}</p>
</body>
</html>
`