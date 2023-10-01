package cmc

import "time"

type ToolsPriceConversion struct {
    Data struct {
        Symbol      string    `json:"symbol"`
        Id          string    `json:"id"`
        Name        string    `json:"name"`
        Amount      int       `json:"amount"`
        LastUpdated time.Time `json:"last_updated"`
        Quote       map[string]struct {
            Price       float64   `json:"price"`
            LastUpdated time.Time `json:"last_updated"`
        } `json:"quote"`
    } `json:"data"`
    Status Status `json:"status"`
}
