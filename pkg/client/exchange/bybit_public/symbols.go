package bybit_public

import (
    "encoding/json"
    "net/http"
    "net/url"
)

func (c *Configuration) GetSymbols() (*SymbolsResponse, error) {
    requestURL := "/v2/public/symbols"

    response, err := request(http.MethodGet, requestURL, url.Values{}, c)
    if err != nil {
        return nil, err
    }

    var symbolsResponse SymbolsResponse

    unmarshalError := json.Unmarshal(response.Body, &symbolsResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &symbolsResponse, nil
}

type SymbolsResponse struct {
    RetCode int      `json:"ret_code"`
    RetMsg  string   `json:"ret_msg"`
    Result  []Symbol `json:"result"`
    ExtCode string   `json:"ext_code"`
    ExtInfo string   `json:"ext_info"`
    TimeNow string   `json:"time_now"`
}

type Symbol struct {
    Name            string `json:"name"`
    Alias           string `json:"alias"`
    Status          string `json:"status"`
    BaseCurrency    string `json:"base_currency"`
    QuoteCurrency   string `json:"quote_currency"`
    PriceScale      int    `json:"price_scale"`
    TakerFee        string `json:"taker_fee"`
    MakerFee        string `json:"maker_fee"`
    FundingInterval int    `json:"funding_interval"`
    LeverageFilter  struct {
        MinLeverage  int    `json:"min_leverage"`
        MaxLeverage  int    `json:"max_leverage"`
        LeverageStep string `json:"leverage_step"`
    } `json:"leverage_filter"`
    PriceFilter struct {
        MinPrice string `json:"min_price"`
        MaxPrice string `json:"max_price"`
        TickSize string `json:"tick_size"`
    } `json:"price_filter"`
    LotSizeFilter struct {
        MaxTradingQty         float64 `json:"max_trading_qty"`
        MinTradingQty         float64 `json:"min_trading_qty"`
        QtyStep               float64 `json:"qty_step"`
        PostOnlyMaxTradingQty string  `json:"post_only_max_trading_qty"`
    } `json:"lot_size_filter"`
}
