package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/vivekprm/go_web_app/viewmodels"
)

func main() {
    templates := populateTemplates()

    http.HandleFunc("/",
    func(w http.ResponseWriter, req *http.Request) {
        requestedFile := req.URL.Path[1:]
        template :=
        templates.Lookup(requestedFile + ".html")

        var context interface{} = nil
        switch requestedFile {
            case "home":
            context = viewmodels.GetHome()
        }
        if template != nil {
            template.Execute(w, context)
        } else {
            w.WriteHeader(404)
        }
    })

    http.HandleFunc("/img/", serveResource)
    http.HandleFunc("/css/", serveResource)
    http.HandleFunc("/scripts/", serveResource)
    http.ListenAndServe(":8080", nil)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
    path := "public" + req.URL.Path
    var contentType string
    if strings.HasSuffix(path, ".css") {
        contentType = "text/css"
    } else if strings.HasSuffix(path, ".png") {
        contentType = "image/png"
    } else if strings.HasSuffix(path, ".jpg") {
        contentType = "image/jpg"
    } else if strings.HasSuffix(path, ".svg") {
        contentType = "image/svg+xml"
    } else if strings.HasSuffix(path, ".js") {
        contentType = "application/javascript"
    } else {
        contentType = "text/plain"
    }

    log.Println(path)
    log.Println(contentType)

    f, err := os.Open(path)

    if err == nil {
        defer f.Close()
        w.Header().Add("Content Type", contentType)
        br := bufio.NewReader(f)
        br.WriteTo(w)
    } else {
        w.WriteHeader(404)
    }
}

func populateTemplates() *template.Template {
    result := template.New("templates")

    basePath := "templates"
    templateFolder, _ := os.Open(basePath)
    defer templateFolder.Close()

    templatePathsRaw, _ := templateFolder.Readdir(-1)
    templatePaths := new([]string)
    for _, pathInfo := range templatePathsRaw {
        log.Println(pathInfo.Name())
        if !pathInfo.IsDir() {
            *templatePaths = append(*templatePaths,
            basePath + "/" + pathInfo.Name())
        }
    }

    result.ParseFiles(*templatePaths...)

    return result
}
