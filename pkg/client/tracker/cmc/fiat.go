package cmc

import (
    "encoding/json"
    "net/url"
)

func (c *Configuration) GetFiatMap() (*Fiat, error) {
    requestURL := "/v1/fiat/map"

    resBody, err := request(requestURL, url.Values{}, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse Fiat
    err = json.Unmarshal(resBody.Body, &cmcResponse)
    if err != nil {
        return nil, err
    }

    return &cmcResponse, nil
}

type Fiat struct {
    Data []struct {
        Id     int    `json:"id"`
        Name   string `json:"name"`
        Sign   string `json:"sign"`
        Symbol string `json:"symbol"`
    } `json:"data"`
    Status Status `json:"status"`
}
