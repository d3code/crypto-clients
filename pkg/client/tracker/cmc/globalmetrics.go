package cmc

import (
    "encoding/json"
    "net/url"
)

func (c *Configuration) GetGlobalMetricsQuotesHistorical(q url.Values) (*GlobalMetricsQuotesHistorical, error) {
    requestURL := "/v1/global-metrics/quotes/historical"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse GlobalMetricsQuotesHistorical

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetGlobalMetricsQuotesLatest(q url.Values) (*GlobalMetricsQuotesLatest, error) {
    requestURL := "/v1/global-metrics/quotes/latest"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse GlobalMetricsQuotesLatest

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}
