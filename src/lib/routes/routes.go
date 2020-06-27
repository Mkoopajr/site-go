package routes

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"

    "github.com/aymerick/raymond"
)

type TemplateHeader struct {
    template *raymond.Template
    context map[string]string
}

func MainTemplate() *TemplateHeader {
    tpl, err := raymond.ParseFile("./views/layouts/main.handlebars")
    if err != nil {
        panic(err)
    }

    tpl.RegisterHelper("writeCSS", func() raymond.SafeString {
        return raymond.SafeString("<link rel='stylesheet' href='/static/css/bundle.min.css'>")
    })

    err = tpl.RegisterPartialFile("./views/partials/nav.handlebars", "nav")
    if err != nil {
        panic(err)
    }

    banImgs, err := ioutil.ReadDir("./static/img/banner")
    if err != nil {
        panic(err)
    }

    var banTpl strings.Builder
    for i, file := range banImgs {
        if i == 0 {
            fmt.Fprintf(&banTpl, "<img src='./static/img/banner/%v' class='banner-img show'>", file.Name())
        } else {
            fmt.Fprintf(&banTpl, "<img src='./static/img/banner/%v' class='banner-img'>", file.Name())
        }
    }

    ctx := map[string]string{
        "banners": banTpl.String(),
    }

    template := TemplateHeader{tpl, ctx}
    return &template
}

func Home(w http.ResponseWriter, r *http.Request) {
    header := MainTemplate()

    content, err := ioutil.ReadFile("./views/home.handlebars")
    if err != nil {
        panic(err)
    }

    header.context["content"] = string(content)

    result, err := header.template.Exec(header.context)
    if err != nil {
        panic(err)
    }

    fmt.Fprint(w, result)
}

func About(w http.ResponseWriter, r *http.Request) {
    header := MainTemplate()

    content, err := ioutil.ReadFile("./views/about.handlebars")
    if err != nil {
        panic(err)
    }

    header.context["content"] = string(content)

    result, err := header.template.Exec(header.context)
    if err != nil {
        panic(err)
    }

    fmt.Fprint(w, result)
}

func Donate(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://paypal.me/HackSIOrg", 301)
}

func Sponsors(w http.ResponseWriter, r *http.Request) {
    header := MainTemplate()

    content, err := ioutil.ReadFile("./views/sponsors.handlebars")
    if err != nil {
        panic(err)
    }

    header.context["content"] = string(content)

    result, err := header.template.Exec(header.context)
    if err != nil {
        panic(err)
    }

    fmt.Fprint(w, result)
}
