package routes

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
    "lib/press"

    "github.com/aymerick/raymond"
)

type TemplateHeader struct {
    template *raymond.Template
    context map[string]interface{}
}

type Template struct {
    Layout *TemplateHeader
    View string
}

type Url struct {
    Destination string
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

    ctx := map[string]interface{}{
        "banners": banTpl.String(),
    }

    template := TemplateHeader{tpl, ctx}
    return &template
}

func (t Template) ServeView(w http.ResponseWriter, r *http.Request) {
    view := fmt.Sprintf("./views/%v.handlebars", t.View)
    content, err := raymond.ParseFile(view)
    if err != nil {
        panic(err)
    }

    var contentString string
    if t.View == "press" {
        var press *press.Press
        press = press.GetPress()
        contentString, err = content.Exec(map[string]interface{}{
            "press": press.Data,
        })
        if err != nil {
            panic(err)
        }
    } else {
        contentString, err = content.Exec(map[string]string{})
    }

    t.Layout.context["content"] = contentString
    t.Layout.context["title"] = t.View
    t.Layout.context["currentYear"] = fmt.Sprintf("%v", time.Now().Year())

    result, err := t.Layout.template.Exec(t.Layout.context)
    if err != nil {
        panic(err)
    }

    fmt.Fprint(w, result)
}

func (url Url) Redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, url.Destination, 301)
}
