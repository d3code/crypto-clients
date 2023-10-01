package bybit

import (
    "encoding/json"
    "net/http"
    "net/url"
)

func (c *Configuration) GetAccountTransferInternal() (*InternalTransferResponse, error) {
    requestURL := "/v5/asset/transfer/query-inter-transfer-list"

    values := url.Values{}

    response, err := request(c, http.MethodGet, requestURL, values, 0)
    if err != nil {
        return nil, err
    }

    var responseObject InternalTransferResponse

    unmarshalError := json.Unmarshal(response.Body, &responseObject)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &responseObject, nil
}

type InternalTransferResponse struct {
    ByBitResponse
    Result struct {
        List           []InternalTransfer `json:"list"`
        NextPageCursor string             `json:"nextPageCursor"`
    } `json:"result"`
}

type InternalTransfer struct {
    TransferId      string `json:"transferId"`
    Coin            string `json:"coin"`
    Amount          string `json:"amount"`
    FromAccountType string `json:"fromAccountType"`
    ToAccountType   string `json:"toAccountType"`
    Timestamp       string `json:"timestamp"`
    Status          string `json:"status"`
}
