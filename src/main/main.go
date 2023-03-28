package main

import (
	"net/http"
	"text/template"
)

type Context struct {
    FirstName string
    Message string
    URL string
    Beers []string
    Title string

}

func main() {
    http.HandleFunc("/", myHandlerFunc)
    http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    templates := template.New("template")
    templates.New("test").Parse(doc)
    templates.New("header").Parse(head)
    templates.New("footer").Parse(foot)
    context := Context{
        "Todd",
        "more beer, please",
        req.URL.Path,
        []string{"New Belgium", "La Fin Du Monde", "The Alchemist"},
        "Favorite Beers",
    }
    templates.Lookup("test").Execute(w, context)
}

const doc = `
{{template "header" .Title}}
<body>
    <h1>{{.FirstName}} says, "{{.Message}}"</h1>
    {{if eq .URL "/nobeer"}}
        <h2>We're out of beer, {{.FirstName}}. Sorry!</h2>
    {{else}}
        <h2>Yes, grab another beer, {{.FirstName}}</h2>
        <ul>
            {{range .Beers}}
            <li>{{.}}</li>
            {{end}}
        </ul>
    {{end}}
    <hr>
    <h2>Here's all the data:</h2>
    <p>{{.}}</p>
</body>
{{template "footer"}}
`

const head = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>{{.}}</title>
</head>
`

const foot = `
</html>
`