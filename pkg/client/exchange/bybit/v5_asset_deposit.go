package bybit

import (
    "encoding/json"
    "net/http"
    "net/url"
)

func (c *Configuration) GetAccountDeposit() (*DepositsResponse, error) {
    requestURL := "/v5/asset/deposit/query-record"

    values := url.Values{}

    response, err := request(c, http.MethodGet, requestURL, values, 0)
    if err != nil {
        return nil, err
    }

    var responseObject DepositsResponse

    unmarshalError := json.Unmarshal(response.Body, &responseObject)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &responseObject, nil
}

type DepositsResponse struct {
    ByBitResponse
    Result struct {
        Rows           []Deposit `json:"rows"`
        NextPageCursor string    `json:"nextPageCursor"`
    } `json:"result"`
}

type Deposit struct {
    Coin          string `json:"coin"`
    Chain         string `json:"chain"`
    Amount        string `json:"amount"`
    TxID          string `json:"txID"`
    Status        int    `json:"status"`
    ToAddress     string `json:"toAddress"`
    Tag           string `json:"tag"`
    DepositFee    string `json:"depositFee"`
    SuccessAt     string `json:"successAt"`
    Confirmations string `json:"confirmations"`
    TxIndex       string `json:"txIndex"`
    BlockHash     string `json:"blockHash"`
}
