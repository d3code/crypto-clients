package binance

import (
    "encoding/json"
    "fmt"
    "github.com/d3code/zlog"
    "net/http"
    "net/url"
)

// GetAccountInfo Get current account information
func (c *Configuration) GetAccountInfo() (*AccountInfo, error) {
    requestURL := "/api/v3/account"

    q := url.Values{}
    response, err := request(http.MethodGet, requestURL, q, c)
    if err != nil {
        return nil, err
    }

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
    }

    var cmcResponse AccountInfo
    unmarshalError := json.Unmarshal(response.Body, &cmcResponse)
    if unmarshalError != nil {
        zlog.Log.Error(unmarshalError)
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

type AccountInfo struct {
    MakerCommission  int `json:"makerCommission"`
    TakerCommission  int `json:"takerCommission"`
    BuyerCommission  int `json:"buyerCommission"`
    SellerCommission int `json:"sellerCommission"`
    CommissionRates  struct {
        Maker  string `json:"maker"`
        Taker  string `json:"taker"`
        Buyer  string `json:"buyer"`
        Seller string `json:"seller"`
    } `json:"commissionRates"`
    CanTrade                   bool      `json:"canTrade"`
    CanWithdraw                bool      `json:"canWithdraw"`
    CanDeposit                 bool      `json:"canDeposit"`
    Brokered                   bool      `json:"brokered"`
    RequireSelfTradePrevention bool      `json:"requireSelfTradePrevention"`
    UpdateTime                 int64     `json:"updateTime"`
    AccountType                string    `json:"accountType"`
    Balances                   []Balance `json:"balances"`
    Permissions                []string  `json:"permissions"`
}

type Balance struct {
    Asset  string `json:"asset"`
    Free   string `json:"free"`
    Locked string `json:"locked"`
}
