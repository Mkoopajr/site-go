package main

import (
    "net/http"
    "lib/routes"

    "github.com/gorilla/mux"
)

func main() {
    mainView := routes.MainTemplate()

    r := mux.NewRouter()
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    r.HandleFunc("/", routes.Template{mainView, "home"}.ServeView)
    r.HandleFunc("/about", routes.Template{mainView, "about"}.ServeView)
    r.HandleFunc("/donate", routes.Donate)
    r.HandleFunc("/sponsors", routes.Template{mainView, "sponsors"}.ServeView)
    r.HandleFunc("/contact", routes.Template{mainView, "contact"}.ServeView)
    r.HandleFunc("/press", routes.Template{mainView, "press"}.ServeView)

    http.Handle("/", r)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    } 
}
