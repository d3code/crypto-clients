package cmc

type Fiat struct {
    Data []struct {
        Id     int    `json:"id"`
        Name   string `json:"name"`
        Sign   string `json:"sign"`
        Symbol string `json:"symbol"`
    } `json:"data"`
    Status Status `json:"status"`
}
