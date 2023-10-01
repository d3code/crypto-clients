package bybit

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
)

// GetOrderHistory returns the closed PnL for a given symbol
//
// category: [spot, linear, inverse, option]
func (c *Configuration) GetOrderHistory(category string) (*[]OrderHistory, error) {
    return c.getOrderHistory(category, "")
}

func (c *Configuration) getOrderHistory(category string, pageCursor string) (*[]OrderHistory, error) {
    requestURL := "/v5/order/history"

    var orderHistory []OrderHistory

    values := url.Values{}
    values.Add("category", category)
    values.Add("cursor", pageCursor)

    response, _ := request(c, http.MethodGet, requestURL, values, 1)
    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("bybit_v5.GetOrderHistory: %v", response.StatusCode)
    }

    var responseObject OrderHistoryResponse
    unmarshalError := json.Unmarshal(response.Body, &responseObject)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    if responseObject.RetCode != 0 {
        return nil, fmt.Errorf("%s - %s", responseObject.RetMsg, responseObject.RetMsg)
    }

    orderHistory = append(orderHistory, responseObject.Result.List...)

    if responseObject.Result.NextPageCursor != "" {
        oh, err := c.getOrderHistory(category, responseObject.Result.NextPageCursor)
        if err != nil {
            return nil, err
        }
        orderHistory = append(orderHistory, *oh...)
    }

    return &orderHistory, nil
}

type OrderHistoryResponse struct {
    ByBitResponse
    Result struct {
        NextPageCursor string         `json:"nextPageCursor"`
        Category       string         `json:"category"`
        List           []OrderHistory `json:"list"`
    } `json:"result"`
}

type OrderHistory struct {
    Symbol             string `json:"symbol"`
    OrderType          string `json:"orderType"`
    OrderLinkId        string `json:"orderLinkId"`
    OrderId            string `json:"orderId"`
    CancelType         string `json:"cancelType"`
    AvgPrice           string `json:"avgPrice"`
    StopOrderType      string `json:"stopOrderType"`
    LastPriceOnCreated string `json:"lastPriceOnCreated"`
    OrderStatus        string `json:"orderStatus"`
    TakeProfit         string `json:"takeProfit"`
    CumExecValue       string `json:"cumExecValue"`
    TriggerDirection   int    `json:"triggerDirection"`
    BlockTradeId       string `json:"blockTradeId"`
    RejectReason       string `json:"rejectReason"`
    IsLeverage         string `json:"isLeverage"`
    Price              string `json:"price"`
    OrderIv            string `json:"orderIv"`
    CreatedTime        string `json:"createdTime"`
    TpTriggerBy        string `json:"tpTriggerBy"`
    PositionIdx        int    `json:"positionIdx"`
    TimeInForce        string `json:"timeInForce"`
    LeavesValue        string `json:"leavesValue"`
    UpdatedTime        string `json:"updatedTime"`
    Side               string `json:"side"`
    TriggerPrice       string `json:"triggerPrice"`
    CumExecFee         string `json:"cumExecFee"`
    SlTriggerBy        string `json:"slTriggerBy"`
    LeavesQty          string `json:"leavesQty"`
    CloseOnTrigger     bool   `json:"closeOnTrigger"`
    CumExecQty         string `json:"cumExecQty"`
    ReduceOnly         bool   `json:"reduceOnly"`
    Qty                string `json:"qty"`
    StopLoss           string `json:"stopLoss"`
    TriggerBy          string `json:"triggerBy"`
}
