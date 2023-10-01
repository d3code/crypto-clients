package bybit

import (
    "encoding/json"
    "net/http"
    "net/url"
)

func (c *Configuration) GetAccountCoinInfo() (*CoinInfoResponse, error) {
    requestURL := "/v5/asset/coin/query-info"

    values := url.Values{}

    response, err := request(c, http.MethodGet, requestURL, values, 0)
    if err != nil {
        return nil, err
    }

    var responseObject CoinInfoResponse
    unmarshalError := json.Unmarshal(response.Body, &responseObject)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &responseObject, nil
}

type CoinInfoResponse struct {
    ByBitResponse
    Result struct {
        Rows []CoinInfo `json:"rows"`
    } `json:"result"`
}

type CoinInfo struct {
    Name         string          `json:"name"`
    Coin         string          `json:"coin"`
    RemainAmount string          `json:"remainAmount"`
    Chains       []CoinInfoChain `json:"chains"`
}

type CoinInfoChain struct {
    ChainType     string `json:"chainType"`
    Confirmation  string `json:"confirmation"`
    WithdrawFee   string `json:"withdrawFee"`
    DepositMin    string `json:"depositMin"`
    WithdrawMin   string `json:"withdrawMin"`
    Chain         string `json:"chain"`
    ChainDeposit  string `json:"chainDeposit"`
    ChainWithdraw string `json:"chainWithdraw"`
    MinAccuracy   string `json:"minAccuracy"`
}
