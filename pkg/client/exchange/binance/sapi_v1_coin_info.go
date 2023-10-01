package binance

import (
    "encoding/json"
    "net/http"
    "net/url"
)

// GetCoinInfo Get information of coins (available for deposit and withdraw) for user
func (c *Configuration) GetCoinInfo() (*[]CoinInfo, error) {
    requestURL := "/sapi/v1/capital/config/getall"

    q := url.Values{}
    response, err := request(http.MethodGet, requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse []CoinInfo

    unmarshalError := json.Unmarshal(response.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

type CoinInfo struct {
    Coin              string        `json:"coin"`
    DepositAllEnable  bool          `json:"depositAllEnable"`
    Free              string        `json:"free"`
    Freeze            string        `json:"freeze"`
    Ipoable           string        `json:"ipoable"`
    Ipoing            string        `json:"ipoing"`
    IsLegalMoney      bool          `json:"isLegalMoney"`
    Locked            string        `json:"locked"`
    Name              string        `json:"name"`
    NetworkList       []NetworkList `json:"networkList"`
    Storage           string        `json:"storage"`
    Trading           bool          `json:"trading"`
    WithdrawAllEnable bool          `json:"withdrawAllEnable"`
    Withdrawing       string        `json:"withdrawing"`
}

type NetworkList struct {
    AddressRegex            string `json:"addressRegex"`
    Coin                    string `json:"coin"`
    DepositDesc             string `json:"depositDesc,omitempty"`
    DepositEnable           bool   `json:"depositEnable"`
    IsDefault               bool   `json:"isDefault"`
    MemoRegex               string `json:"memoRegex"`
    MinConfirm              int    `json:"minConfirm"`
    Name                    string `json:"name"`
    Network                 string `json:"network"`
    ResetAddressStatus      bool   `json:"resetAddressStatus"`
    SpecialTips             string `json:"specialTips"`
    UnLockConfirm           int    `json:"unLockConfirm"`
    WithdrawDesc            string `json:"withdrawDesc,omitempty"`
    WithdrawEnable          bool   `json:"withdrawEnable"`
    WithdrawFee             string `json:"withdrawFee"`
    WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
    WithdrawMax             string `json:"withdrawMax"`
    WithdrawMin             string `json:"withdrawMin"`
    SameAddress             bool   `json:"sameAddress"`
}
