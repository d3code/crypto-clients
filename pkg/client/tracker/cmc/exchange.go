package cmc

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/zlog"
    "net/url"
    "strings"
)

func (c *Configuration) GetExchangeInfo(id string) (*ExchangeInfo, error) {
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
        var cmcResponse ExchangeInfo
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

    var cmcResponse ExchangeInfo
    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

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

func (c *Configuration) GetExchangeListingsLatest(q url.Values) (*ExchangeListingsLatest, error) {
    requestURL := "/v1/exchange/listings/latest"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }
    var cmcResponse ExchangeListingsLatest

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetExchangeMarketPairsLatest(q url.Values) (*ExchangeMarketPairsLatest, error) {
    requestURL := "/v1/exchange/market-pairs/latest"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }
    var cmcResponse ExchangeMarketPairsLatest

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetExchangeQuotesHistorical(q url.Values) (*ExchangeQuotesHistorical, error) {
    requestURL := "/v1/exchange/quotes/historical"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }
    var cmcResponse ExchangeQuotesHistorical

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetExchangeQuotesLatest(q url.Values) (*ExchangeQuotesLatest, error) {
    requestURL := "/v1/exchange/quotes/latest"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }
    var cmcResponse ExchangeQuotesLatest

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}
