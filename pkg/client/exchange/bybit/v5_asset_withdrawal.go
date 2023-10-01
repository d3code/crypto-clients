package bybit

import (
    "encoding/json"
    "net/http"
    "net/url"
)

func (c *Configuration) GetAccountWithdrawal() (*WithdrawalsResponse, error) {
    requestURL := "/v5/asset/withdraw/query-record"

    values := url.Values{}

    response, err := request(c, http.MethodGet, requestURL, values, 0)
    if err != nil {
        return nil, err
    }

    var responseObject WithdrawalsResponse

    unmarshalError := json.Unmarshal(response.Body, &responseObject)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &responseObject, nil
}

type WithdrawalsResponse struct {
    ByBitResponse
    Result struct {
        Rows           []Withdrawal `json:"rows"`
        NextPageCursor string       `json:"nextPageCursor"`
    } `json:"result"`
}

type Withdrawal struct {
    Coin         string `json:"coin"`
    Chain        string `json:"chain"`
    Amount       string `json:"amount"`
    TxID         string `json:"txID"`
    Status       string `json:"status"`
    ToAddress    string `json:"toAddress"`
    Tag          string `json:"tag"`
    WithdrawFee  string `json:"withdrawFee"`
    CreateTime   string `json:"createTime"`
    UpdateTime   string `json:"updateTime"`
    WithdrawId   string `json:"withdrawId"`
    WithdrawType int    `json:"withdrawType"`
}
