package bybit

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/zlog"
    "net/http"
    "net/url"
)

// GetExecutionHistory
// Query users' execution records, sorted by execTime in descending order
// However, for Normal spot, they are sorted by execId in descending order
//
// category: [spot, linear, inverse, option]
func (c *Configuration) GetExecutionHistory(category string) (*[]ExecutionHistory, error) {
    return c.getExecutionHistory(category, "")
}

func (c *Configuration) getExecutionHistory(category string, pageCursor string) (*[]ExecutionHistory, error) {
    requestURL := "/v5/execution/list"

    var orderHistory []ExecutionHistory

    values := url.Values{}
    values.Add("category", category)
    values.Add("cursor", pageCursor)

    response, _ := request(c, http.MethodGet, requestURL, values, 1)
    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("bybit_v5.GetOrderHistory: %v", response.StatusCode)
    }

    var responseObject ExecutionHistoryResponse
    unmarshalError := json.Unmarshal(response.Body, &responseObject)
    if unmarshalError != nil {
        zlog.Log.Error(string(response.Body))
        return nil, unmarshalError
    }

    if responseObject.RetCode != 0 {
        return nil, fmt.Errorf("%s - %s", responseObject.RetMsg, responseObject.RetMsg)
    }

    orderHistory = append(orderHistory, responseObject.Result.List...)

    if responseObject.Result.NextPageCursor != "" {
        oh, err := c.getExecutionHistory(category, responseObject.Result.NextPageCursor)
        if err != nil {
            return nil, err
        }
        orderHistory = append(orderHistory, *oh...)
    }

    return &orderHistory, nil
}

type ExecutionHistoryResponse struct {
    ByBitResponse
    Result struct {
        NextPageCursor string             `json:"nextPageCursor"`
        Category       string             `json:"category"`
        List           []ExecutionHistory `json:"list"`
    } `json:"result"`
}

type ExecutionHistory struct {
    Symbol          string `json:"symbol"`
    OrderType       string `json:"orderType"`
    UnderlyingPrice string `json:"underlyingPrice"`
    OrderLinkId     string `json:"orderLinkId"`
    Side            string `json:"side"`
    IndexPrice      string `json:"indexPrice"`
    OrderId         string `json:"orderId"`
    StopOrderType   string `json:"stopOrderType"`
    LeavesQty       string `json:"leavesQty"`
    ExecTime        string `json:"execTime"`
    IsMaker         bool   `json:"isMaker"`
    ExecFee         string `json:"execFee"`
    FeeRate         string `json:"feeRate"`
    ExecId          string `json:"execId"`
    TradeIv         string `json:"tradeIv"`
    BlockTradeId    string `json:"blockTradeId"`
    MarkPrice       string `json:"markPrice"`
    ExecPrice       string `json:"execPrice"`
    MarkIv          string `json:"markIv"`
    OrderQty        string `json:"orderQty"`
    OrderPrice      string `json:"orderPrice"`
    ExecValue       string `json:"execValue"`
    ExecType        string `json:"execType"`
    ExecQty         string `json:"execQty"`
}
