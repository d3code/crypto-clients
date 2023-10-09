package asx

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
)

func (c *Configuration) Details(symbol string) (*DetailsResponse, error) {
    requestURL := fmt.Sprintf("/asx-research/1.0/companies/%s/about", symbol)

    values := url.Values{}
    response, err := request(http.MethodGet, requestURL, values, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse DetailsResponse
    unmarshalError := json.Unmarshal(response.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

type DetailsResponse struct {
    Data Detail `json:"data"`
}

type Detail struct {
    DisplayName          string        `json:"displayName"`
    IssueType            string        `json:"issueType"`
    Symbol               string        `json:"symbol"`
    Xid                  string        `json:"xid"`
    AddressContact       Address       `json:"addressContact"`
    AddressShareRegistry ShareRegistry `json:"addressShareRegistry"`
    Description          string        `json:"description"`
    Directors            []Director    `json:"directors"`
    ForeignExempt        bool          `json:"foreignExempt"`
    WebsiteUrl           string        `json:"websiteUrl"`
    Secretaries          []Secretary   `json:"secretaries"`
}

type Address struct {
    Address string `json:"address"`
    Fax     string `json:"fax"`
    Phone   string `json:"phone"`
}

type ShareRegistry struct {
    Address   string `json:"address"`
    Attention string `json:"attention"`
    Phone     string `json:"phone"`
}

type Director struct {
    Name  string `json:"name"`
    Title string `json:"title"`
}

type Secretary struct {
    Name  string `json:"name"`
    Title string `json:"title"`
}
