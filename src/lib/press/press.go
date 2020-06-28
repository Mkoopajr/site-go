package press

import (
    "io/ioutil"
    "encoding/json"
)

type Press struct {
    Data []Story `json:"data"`
}

type Story struct {
    Title string `json:"title"`
    Content string `json:"content"`
}

func (p *Press)  GetPress() *Press {
    var press Press

    pressBytes, err := ioutil.ReadFile("./views/data/news.json")
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(pressBytes, &press)
    if err != nil {
        panic(err)
    }

    return &press
}
