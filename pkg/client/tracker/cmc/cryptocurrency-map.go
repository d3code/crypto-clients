package cmc

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/zlog"
    "net/url"
    "time"
)

func (c *Configuration) CryptocurrencyMap() []*CryptocurrencyMapResponse {
    var responses []*CryptocurrencyMapResponse

    responses = append(responses, cryptocurrencyMapPage(c, 1)...)
    return responses
}

func cryptocurrencyMapPage(c *Configuration, start int) []*CryptocurrencyMapResponse {
    requestURL := "/v1/cryptocurrency/map"

    values := url.Values{}
    values.Set("listing_status", "active,inactive,untracked")
    values.Set("start", fmt.Sprintf("%d", start))

    response, err := request(requestURL, values, c)
    if err != nil {
        zlog.Log.Error(err)
    }

    var responseMap []*CryptocurrencyMapResponse

    var cmcResponse CryptocurrencyMapResponse
    err = json.Unmarshal(response.Body, &cmcResponse)
    if err != nil {
        zlog.Log.Error(err)
    }

    responseMap = append(responseMap, &cmcResponse)
    if len(cmcResponse.Data) == 10000 {
        re := cryptocurrencyMapPage(c, start+10000)
        responseMap = append(responseMap, re...)
    }

    return responseMap
}

type Cryptocurrency struct {
    Id                  int       `json:"id"`
    Rank                *int      `json:"rank"`
    Name                string    `json:"name"`
    Symbol              string    `json:"symbol"`
    Slug                string    `json:"slug"`
    IsActive            int       `json:"is_active"`
    FirstHistoricalData time.Time `json:"first_historical_data"`
    LastHistoricalData  time.Time `json:"last_historical_data"`
    Platform            *Platform `json:"platform"`
}

type Platform struct {
    Id           int    `json:"id"`
    Name         string `json:"name"`
    Symbol       string `json:"symbol"`
    Slug         string `json:"slug"`
    TokenAddress string `json:"token_address"`
}

type CryptocurrencyMapResponse struct {
    Data   []Cryptocurrency `json:"data"`
    Status Status           `json:"status"`
}
