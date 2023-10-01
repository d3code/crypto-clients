package coin_gecko

import (
    "encoding/json"
    "net/url"
)

func (c *Configuration) CoinsList() (*[]CoinList, error) {
    requestURL := "/api/v3/coins/list"

    requestParams := url.Values{}
    requestParams.Set("include_platform", "true")

    resBody, err := doGetRequest(requestURL, requestParams, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse []CoinList

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

type CoinList struct {
    Id        string            `json:"id"`
    Symbol    string            `json:"symbol"`
    Name      string            `json:"name"`
    Platforms map[string]string `json:"platforms"`
}
