package cmc

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/zlog"
    "net/url"
    "time"
)

func (c *Configuration) CryptocurrencyInfo(id string) (*CryptocurrencyInfoResponse, error) {
    requestURL := "/v2/cryptocurrency/info"

    values := url.Values{}
    values.Add("id", id)
    values.Add("skip_invalid", "true")

    response, err := request(requestURL, values, c)
    if err != nil {
        return nil, err
    }

    if response.StatusCode == 429 {
        zlog.Log.Warn("Rate limit exceeded, sleeping for 62 seconds")

        time.Sleep(62 * time.Second)
        response, err = request(requestURL, values, c)
    }

    if response.StatusCode != 200 {
        return nil, fmt.Errorf("status code: %v %s", response.StatusCode, string(response.Body))
    }

    var cmcResponse CryptocurrencyInfoResponse
    unmarshalError := json.Unmarshal(response.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

type CryptocurrencyInfoResponse struct {
    Data   map[string]CryptocurrencyInfo `json:"data"`
    Status Status                        `json:"status"`
}

type CryptocurrencyInfo struct {
    Logo                          string              `json:"logo"`
    Id                            int                 `json:"id"`
    Name                          string              `json:"name"`
    Symbol                        string              `json:"symbol"`
    Slug                          string              `json:"slug"`
    Description                   string              `json:"description"`
    DateAdded                     time.Time           `json:"date_added"`
    DateLaunched                  time.Time           `json:"date_launched"`
    Tags                          []string            `json:"tags"`
    Category                      string              `json:"category"`
    Notice                        string              `json:"notice"`
    SelfReportedCirculatingSupply *float32            `json:"self_reported_circulating_supply"`
    SelfReportedMarketCap         *float32            `json:"self_reported_market_cap"`
    SelfReportedTags              *[]string           `json:"self_reported_tags"`
    InfiniteSupply                bool                `json:"infinite_supply"`
    ContractAddress               *[]ContractAddress  `json:"contract_address"`
    Platform                      *PlatformInfo       `json:"platform"`
    Urls                          map[string][]string `json:"urls"`
}

type Coin struct {
    Id     string `json:"id"`
    Name   string `json:"name"`
    Symbol string `json:"symbol"`
    Slug   string `json:"slug"`
}

type ContractPlatform struct {
    Name string `json:"name"`
    Coin Coin   `json:"coin"`
}

type ContractAddress struct {
    ContractAddress string           `json:"contract_address"`
    Platform        ContractPlatform `json:"platform"`
}

type PlatformInfo struct {
    Id           string `json:"id"`
    Name         string `json:"name"`
    Symbol       string `json:"symbol"`
    Slug         string `json:"slug"`
    TokenAddress string `json:"token_address"`
}
