package binance

import (
    "encoding/json"
    "net/http"
    "net/url"
)

func (c *Configuration) GetOrders(symbol string) (*[]Order, error) {
    requestURL := "/api/v3/allOrders"

    values := url.Values{}
    values.Set("symbol", symbol)

    response, err := request(http.MethodGet, requestURL, values, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse []Order

    unmarshalError := json.Unmarshal(response.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

type Order struct {
    Symbol              string `json:"symbol"`
    OrderId             int    `json:"orderId"`
    OrderListId         int    `json:"orderListId"`
    ClientOrderId       string `json:"clientOrderId"`
    Price               string `json:"price"`
    OrigQty             string `json:"origQty"`
    ExecutedQty         string `json:"executedQty"`
    CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
    Status              string `json:"status"`
    TimeInForce         string `json:"timeInForce"`
    Type                string `json:"type"`
    Side                string `json:"side"`
    StopPrice           string `json:"stopPrice"`
    IcebergQty          string `json:"icebergQty"`
    Time                int64  `json:"time"`
    UpdateTime          int64  `json:"updateTime"`
    IsWorking           bool   `json:"isWorking"`
    OrigQuoteOrderQty   string `json:"origQuoteOrderQty"`
}
