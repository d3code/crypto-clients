package cmc

import (
    "encoding/json"
    "net/url"
    "time"
)

func (c *Configuration) GetExchanges() (*ExchangeMap, error) {
    requestURL := "/v1/exchange/map"
    values := url.Values{}
    values.Add("listing_status", "active,inactive,untracked")

    resBody, err := request(requestURL, values, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse ExchangeMap

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

type Exchange struct {
    Id                  int       `json:"id"`
    Name                string    `json:"name"`
    Slug                string    `json:"slug"`
    IsActive            int       `json:"is_active"`
    IsListed            int       `json:"is_listed"`
    IsRedistributable   int       `json:"is_redistributable"`
    FirstHistoricalData time.Time `json:"first_historical_data"`
    LastHistoricalData  time.Time `json:"last_historical_data"`
}

type ExchangeMap struct {
    Data   []Exchange `json:"data"`
    Status Status     `json:"status"`
}
