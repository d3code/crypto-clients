package bybit

import (
    "encoding/json"
    "net/http"
    "net/url"
)

func (c *Configuration) GetAccountTransactionLog() (*TransactionLogResponse, error) {
    requestURL := "/v5/account/transaction-log"

    values := url.Values{}

    response, err := request(c, http.MethodGet, requestURL, values, 0)
    if err != nil {
        return nil, err
    }

    var responseObject TransactionLogResponse

    unmarshalError := json.Unmarshal(response.Body, &responseObject)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &responseObject, nil
}

type TransactionLogResponse struct {
    ByBitResponse
    Result struct {
        NextPageCursor string           `json:"nextPageCursor"`
        List           []TransactionLog `json:"list"`
    } `json:"result"`
}

type TransactionLog struct {
    Symbol          string `json:"symbol"`
    Side            string `json:"side"`
    Funding         string `json:"funding"`
    OrderLinkId     string `json:"orderLinkId"`
    OrderId         string `json:"orderId"`
    Fee             string `json:"fee"`
    Change          string `json:"change"`
    CashFlow        string `json:"cashFlow"`
    TransactionTime string `json:"transactionTime"`
    Type            string `json:"type"`
    FeeRate         string `json:"feeRate"`
    BonusChange     string `json:"bonusChange"`
    Size            string `json:"size"`
    Qty             string `json:"qty"`
    CashBalance     string `json:"cashBalance"`
    Currency        string `json:"currency"`
    Category        string `json:"category"`
    TradePrice      string `json:"tradePrice"`
    TradeId         string `json:"tradeId"`
}
