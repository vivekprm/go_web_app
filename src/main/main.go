package main

import (
	"net/http"
	"text/template"
)

func main() {
    http.HandleFunc("/", myHandlerFunc)
    http.ListenAndServe(":8080", nil)
    // nil means use default ServeMux
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        tmpl.Execute(w, req.URL.Path[1:])
        // nil means no data to pass in
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
    <h1>Hello {{.}}</h1>
</body>
</html>
`