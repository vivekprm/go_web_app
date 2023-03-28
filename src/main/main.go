package main

import (
	"net/http"
	//"text/template"
	"html/template"
)

//var Message string = "more beer, please sir"
var Message string = "alert('you have been pwned')"
//var Message string = "<script>alert('you have been pwned, BIATCH')</script>"

func main() {
    http.HandleFunc("/", myHandlerFunc)
    http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        tmpl.Execute(w, Message)
    }
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>Injection Safe</title>
</head>
<body>
    <p>{{.}}</p>
    <script>{{.}}</script>
</body>
</html>
`