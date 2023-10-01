package cmc

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/zlog"
    "net/url"
    "time"
)

func (c *Configuration) GetCryptocurrencyAirdrop(q url.Values) (*CryptocurrencyAirdrop, error) {
    requestURL := "/v1/cryptocurrency/airdrop"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyAirdrop

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyAirdrops(q url.Values) (*CryptocurrencyAirdrops, error) {
    requestURL := "/v1/cryptocurrency/airdrops"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyAirdrops

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyCategories(q url.Values) (*CryptocurrencyCategories, error) {
    requestURL := "/v1/cryptocurrency/categories"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyCategories

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyCategory(q url.Values) (*CryptocurrencyCategory, error) {
    requestURL := "/v1/cryptocurrency/category"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyCategory

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyInfo(q url.Values) (*CryptocurrencyInfo, error) {
    requestURL := "/v1/cryptocurrency/info"

    response, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    if response.StatusCode == 429 {
        zlog.Log.Warn("Rate limit exceeded, sleeping for 62 seconds")

        time.Sleep(62 * time.Second)
        response, err = request(requestURL, q, c)
    }

    if response.StatusCode != 200 {
        return nil, fmt.Errorf("status code: %v %s", response.StatusCode, string(response.Body))
    }

    var cmcResponse CryptocurrencyInfo
    unmarshalError := json.Unmarshal(response.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyListingsHistorical(q url.Values) (*CryptocurrencyListingsHistorical, error) {
    requestURL := "/v1/cryptocurrency/listings/historical"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyListingsHistorical

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyListingsLatest(q url.Values) (*CryptocurrencyListingsLatest, error) {
    requestURL := "/v1/cryptocurrency/listings/latest"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyListingsLatest

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyListingsNew(q url.Values) (*CryptocurrencyListingsNew, error) {
    requestURL := "/v1/cryptocurrency/listings/new"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyListingsNew

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyTrendingGainersLosers(q url.Values) (*CryptocurrencyTrendingGainersLosers, error) {
    requestURL := "/v1/cryptocurrency/trending/gainers-losers"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyTrendingGainersLosers

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyTrendingLatest(q url.Values) (*CryptocurrencyTrendingLatest, error) {
    requestURL := "/v1/cryptocurrency/trending/latest"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyTrendingLatest

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyTrendingMostVisited(q url.Values) (*CryptocurrencyTrendingMostVisited, error) {
    requestURL := "/v1/cryptocurrency/trending/most-visited"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyTrendingMostVisited

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyMarketPairs(q url.Values) (*CryptocurrencyMarketPairsLatest, error) {
    requestURL := "/v2/cryptocurrency/market-pairs/latest"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyMarketPairsLatest

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyOHLCVHistorical(q url.Values) (*CryptocurrencyOHLCVHistorical, error) {
    requestURL := "/v2/cryptocurrency/ohlcv/historical"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyOHLCVHistorical

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyOHLCVLatest(q url.Values) (*CryptocurrencyOHLCVLatest, error) {
    requestURL := "/v2/cryptocurrency/ohlcv/latest"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyOHLCVLatest

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyPricePerformanceStats(q url.Values) (*CryptocurrencyPricePerformanceStats, error) {
    requestURL := "/v2/cryptocurrency/price-performance-stats/latest"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyPricePerformanceStats

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyQuotesHistorical(q url.Values) (*CryptocurrencyQuotesHistorical, error) {
    requestURL := "/v2/cryptocurrency/quotes/historical"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyQuotesHistorical

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

func (c *Configuration) GetCryptocurrencyQuotesLatest(q url.Values) (*CryptocurrencyQuotesLatest, error) {
    requestURL := "/v2/cryptocurrency/quotes/latest"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse CryptocurrencyQuotesLatest

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}
