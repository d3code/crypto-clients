package binance_public

import (
    "encoding/json"
    "net/http"
    "net/url"
)

// GetExchangeInfo Current exchange trading rules and symbol information
func (c *Configuration) GetExchangeInfo() (*ExchangeInfo, error) {
    requestURL := "/api/v3/exchangeInfo"

    q := url.Values{}
    response, err := request(http.MethodGet, requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var exchangeInfo ExchangeInfo
    unmarshalError := json.Unmarshal(response.Body, &exchangeInfo)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &exchangeInfo, nil
}

type ExchangeInfo struct {
    Timezone        string        `json:"timezone"`
    ServerTime      int64         `json:"serverTime"`
    RateLimits      []RateLimit   `json:"rateLimits"`
    ExchangeFilters []interface{} `json:"exchangeFilters"`
    Symbols         []Symbol      `json:"symbols"`
}

type RateLimit struct {
    Type        string `json:"rateLimitType"`
    Interval    string `json:"interval"`
    IntervalNum int    `json:"intervalNum"`
    Limit       int    `json:"limit"`
}

type Symbol struct {
    Symbol                          string   `json:"symbol"`
    Status                          string   `json:"status"`
    BaseAsset                       string   `json:"baseAsset"`
    BaseAssetPrecision              int      `json:"baseAssetPrecision"`
    QuoteAsset                      string   `json:"quoteAsset"`
    QuotePrecision                  int      `json:"quotePrecision"`
    QuoteAssetPrecision             int      `json:"quoteAssetPrecision"`
    BaseCommissionPrecision         int      `json:"baseCommissionPrecision"`
    QuoteCommissionPrecision        int      `json:"quoteCommissionPrecision"`
    OrderTypes                      []string `json:"orderTypes"`
    IcebergAllowed                  bool     `json:"icebergAllowed"`
    OcoAllowed                      bool     `json:"ocoAllowed"`
    QuoteOrderQtyMarketAllowed      bool     `json:"quoteOrderQtyMarketAllowed"`
    AllowTrailingStop               bool     `json:"allowTrailingStop"`
    CancelReplaceAllowed            bool     `json:"cancelReplaceAllowed"`
    IsSpotTradingAllowed            bool     `json:"isSpotTradingAllowed"`
    IsMarginTradingAllowed          bool     `json:"isMarginTradingAllowed"`
    Filters                         []Filter `json:"filters"`
    Permissions                     []string `json:"permissions"`
    DefaultSelfTradePreventionMode  string   `json:"defaultSelfTradePreventionMode"`
    AllowedSelfTradePreventionModes []string `json:"allowedSelfTradePreventionModes"`
}

type Filter struct {
    FilterType            string `json:"filterType"`
    MinPrice              string `json:"minPrice,omitempty"`
    MaxPrice              string `json:"maxPrice,omitempty"`
    TickSize              string `json:"tickSize,omitempty"`
    MinQty                string `json:"minQty,omitempty"`
    MaxQty                string `json:"maxQty,omitempty"`
    StepSize              string `json:"stepSize,omitempty"`
    MinNotional           string `json:"minNotional,omitempty"`
    ApplyToMarket         bool   `json:"applyToMarket,omitempty"`
    AvgPriceMins          int    `json:"avgPriceMins,omitempty"`
    Limit                 int    `json:"limit,omitempty"`
    MinTrailingAboveDelta int    `json:"minTrailingAboveDelta,omitempty"`
    MaxTrailingAboveDelta int    `json:"maxTrailingAboveDelta,omitempty"`
    MinTrailingBelowDelta int    `json:"minTrailingBelowDelta,omitempty"`
    MaxTrailingBelowDelta int    `json:"maxTrailingBelowDelta,omitempty"`
    BidMultiplierUp       string `json:"bidMultiplierUp,omitempty"`
    BidMultiplierDown     string `json:"bidMultiplierDown,omitempty"`
    AskMultiplierUp       string `json:"askMultiplierUp,omitempty"`
    AskMultiplierDown     string `json:"askMultiplierDown,omitempty"`
    MaxNumOrders          int    `json:"maxNumOrders,omitempty"`
    MaxNumAlgoOrders      int    `json:"maxNumAlgoOrders,omitempty"`
    MaxPosition           string `json:"maxPosition,omitempty"`
}
