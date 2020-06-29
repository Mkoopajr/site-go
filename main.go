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
    r.HandleFunc("/sponsors", routes.Template{mainView, "sponsors"}.ServeView)
    r.HandleFunc("/contact", routes.Template{mainView, "contact"}.ServeView)
    r.HandleFunc("/press", routes.Template{mainView, "press"}.ServeView)

    r.HandleFunc("/donate", routes.Url{"https://paypal.me/HackSIOrg"}.Redirect)
    r.HandleFunc("/twitter", routes.Url{"https://twitter.com/hacksi"}.Redirect)
    r.HandleFunc("/instagram", routes.Url{"https://instagram.com/hacksi2019"}.Redirect)
    r.HandleFunc("/github", routes.Url{"https://github.com/HackSI"}.Redirect)
    r.HandleFunc("/facebook", routes.Url{"https://facebook.com/HackSouthernIllinois"}.Redirect)
    r.HandleFunc("/flickr", routes.Url{"https://www.flickr.com/groups/hacksi/pool/"}.Redirect)

    http.Handle("/", r)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    } 
}
