package asx

import (
    "encoding/json"
    "net/http"
    "net/url"
)

func (c *Configuration) Directory() (*DirectoryResponse, error) {
    requestURL := "/asx-research/1.0/companies/directory"

    values := url.Values{}
    values.Add("page", "0")
    values.Add("itemsPerPage", "5000")
    values.Add("sort", "companyName")
    values.Add("order", "ascending")
    values.Add("includeFilterOptions", "true")
    values.Add("recentListingsOnly", "false")

    response, err := request(http.MethodGet, requestURL, values, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse DirectoryResponse
    unmarshalError := json.Unmarshal(response.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

type DirectoryResponse struct {
    Data DirectoryListing `json:"data"`
}

type DirectoryListing struct {
    Items         []Company      `json:"items"`
    Count         int            `json:"count"`
    FilterOptions []FilterOption `json:"filterOptions"`
}

type Company struct {
    Symbol                    string      `json:"symbol"`
    DisplayName               string      `json:"displayName"`
    Industry                  string      `json:"industry"`
    DateListed                string      `json:"dateListed"`
    MarketCap                 interface{} `json:"marketCap"`
    Xid                       string      `json:"xid"`
    PriceChangeFiveDayPercent float64     `json:"priceChangeFiveDayPercent"`
    IsRecentListing           bool        `json:"isRecentListing"`
    StatusCode                string      `json:"statusCode"`
}

type FilterOption struct {
    Input   string   `json:"input"`
    Options []string `json:"options"`
}
