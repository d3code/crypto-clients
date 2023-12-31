package cmc

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/zlog"
    "net/url"
    "strings"
    "time"
)

func (c *Configuration) GetExchangeInfo(id string) (*ExchangeInfoResponse, error) {
    requestURL := "/v1/exchange/info"

    values := url.Values{}
    values.Add("id", id)

    var invalidValues []string
    resBody, err := request(requestURL, values, c)
    if err != nil {
        return nil, err
    }

    if resBody.Body == nil {
        return nil, fmt.Errorf("body is nil")
    }

    if resBody.StatusCode == 400 {
        var cmcResponse ExchangeInfoResponse
        unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
        if unmarshalError != nil {
            return nil, unmarshalError
        }

        if strings.HasPrefix(cmcResponse.Status.ErrorMessage, "Invalid value") {
            invalidValueString := strings.Split(cmcResponse.Status.ErrorMessage, ": ")[1]
            invalidValueString = strings.TrimPrefix(invalidValueString, "\"")
            invalidValueString = strings.TrimSuffix(invalidValueString, "\"")

            invalidValues = strings.Split(invalidValueString, ",")
            zlog.Log.Warnf("Invalid values: %v", invalidValues)
        }
    }

    if resBody.StatusCode != 200 {
        zlog.Log.Error(string(resBody.Body))
        return nil, fmt.Errorf("status code %d", resBody.StatusCode)
    }

    var cmcResponse ExchangeInfoResponse
    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

type ExchangeInfoResponse struct {
    Data   map[string]ExchangeInfo `json:"data"`
    Status Status                  `json:"status"`
}

type ExchangeInfo struct {
    Id                    int                 `json:"id"`
    Name                  string              `json:"name"`
    Slug                  string              `json:"slug"`
    Logo                  string              `json:"logo"`
    Description           *string             `json:"description"`
    DateLaunched          *time.Time          `json:"date_launched"`
    Notice                string              `json:"notice"`
    Countries             []string            `json:"countries"`
    Fiats                 []string            `json:"fiats"`
    Type                  string              `json:"type"`
    MakerFee              float64             `json:"maker_fee"`
    TakerFee              float64             `json:"taker_fee"`
    WeeklyVisits          *int                `json:"weekly_visits"`
    SpotVolumeUsd         *float64            `json:"spot_volume_usd"`
    SpotVolumeLastUpdated *time.Time          `json:"spot_volume_last_updated"`
    PorStatus             *int                `json:"porStatus"`
    PorAuditStatus        *int                `json:"porAuditStatus"`
    Urls                  map[string][]string `json:"urls"`
}
