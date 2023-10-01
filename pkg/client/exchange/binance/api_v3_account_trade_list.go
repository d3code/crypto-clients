package binance

import (
    "encoding/json"
    "net/http"
    "net/url"
)

func (c *Configuration) GetAccountTradeList(symbol string) (*[]AccountTradeList, error) {
    requestURL := "/api/v3/myTrades"

    values := url.Values{}
    values.Set("symbol", symbol)

    response, err := request(http.MethodGet, requestURL, values, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse []AccountTradeList

    unmarshalError := json.Unmarshal(response.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

type AccountTradeList struct {
    Symbol          string `json:"symbol"`
    Id              int    `json:"id"`
    OrderId         int    `json:"orderId"`
    OrderListId     int    `json:"orderListId"`
    Price           string `json:"price"`
    Qty             string `json:"qty"`
    QuoteQty        string `json:"quoteQty"`
    Commission      string `json:"commission"`
    CommissionAsset string `json:"commissionAsset"`
    Time            int64  `json:"time"`
    IsBuyer         bool   `json:"isBuyer"`
    IsMaker         bool   `json:"isMaker"`
    IsBestMatch     bool   `json:"isBestMatch"`
}
