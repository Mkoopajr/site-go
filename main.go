package main

import (
    "net/http"
    "lib/routes"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    r.HandleFunc("/", routes.Home)
    r.HandleFunc("/about", routes.About)
    r.HandleFunc("/donate", routes.Donate)
    r.HandleFunc("/sponsors", routes.Sponsors)

    http.Handle("/", r)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    } 
}
