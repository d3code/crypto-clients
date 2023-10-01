package cmc

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/zlog"
    "net/url"
    "time"
)

func (c *Configuration) CryptocurrencyMap() map[int]Cryptocurrency {
    var responses = make(map[int]Cryptocurrency)

    cryptocurrencyMapPage(c, 1, &responses)
    return responses
}

func cryptocurrencyMapPage(c *Configuration, start int, responses *map[int]Cryptocurrency) {
    requestURL := "/v1/cryptocurrency/map"

    values := url.Values{}
    values.Set("listing_status", "active,inactive,untracked")
    values.Set("start", fmt.Sprintf("%d", start))

    resBody, err := request(requestURL, values, c)
    if err != nil {
        zlog.Log.Error(err)
    }

    var cmcResponse CryptocurrencyMapResponse
    err = json.Unmarshal(resBody.Body, &cmcResponse)
    if err != nil {
        zlog.Log.Error(err)
    }

    for _, cryptocurrency := range cmcResponse.Data {
        (*responses)[cryptocurrency.Id] = cryptocurrency
    }

    if len(cmcResponse.Data) == 10000 {
        cryptocurrencyMapPage(c, start+10000, responses)
    }
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
