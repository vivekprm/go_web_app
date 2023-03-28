package main

import (
	"html/template"
	"net/http"
	"os"
)

func main() {
    templates := populateTemplates()

    http.HandleFunc("/",
    func(w http.ResponseWriter, req *http.Request) {
        requestedFile := req.URL.Path[1:]
        template := templates.Lookup(requestedFile + ".html")

        if template != nil {
            template.Execute(w, nil)
        } else {
            w.WriteHeader(404)
        }
    })


    http.ListenAndServe(":8080", nil)
}

func populateTemplates() *template.Template {
    result := template.New("templates")

    basePath := "templates"
    templateFolder, _ := os.Open(basePath)
    defer templateFolder.Close()

    templatePathsRaw, _ := templateFolder.Readdir(-1)
    // -1 means all of the contents
    templatePaths := new([]string)
    for _, pathInfo := range templatePathsRaw {
        if !pathInfo.IsDir() {
            *templatePaths = append(*templatePaths,
            basePath + "/" + pathInfo.Name())
        }
    }

    result.ParseFiles(*templatePaths...)

    return result
}