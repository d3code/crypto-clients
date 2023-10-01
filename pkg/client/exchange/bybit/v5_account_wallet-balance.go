package bybit

import (
    "encoding/json"
    "net/http"
    "net/url"
)

// GetAccountWalletBalance returns the wallet balance
//
// Account type
// Unified account: UNIFIED (trade spot/linear/options), CONTRACT(trade inverse)
// Normal account: CONTRACT, SPOT
func (c *Configuration) GetAccountWalletBalance(accountType string) (*WalletBalanceResponse, error) {
    requestURL := "/v5/account/wallet-balance"

    values := url.Values{}
    values.Add("accountType", accountType)

    response, err := request(c, http.MethodGet, requestURL, values, 0)
    if err != nil {
        return nil, err
    }

    var responseObject WalletBalanceResponse

    unmarshalError := json.Unmarshal(response.Body, &responseObject)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &responseObject, nil
}

type WalletBalanceResponse struct {
    ByBitResponse
    Result struct {
        List []WalletBalance `json:"list"`
    } `json:"result"`
}

type WalletBalance struct {
    TotalEquity            string              `json:"totalEquity"`
    AccountIMRate          string              `json:"accountIMRate"`
    TotalMarginBalance     string              `json:"totalMarginBalance"`
    TotalInitialMargin     string              `json:"totalInitialMargin"`
    AccountType            string              `json:"accountType"`
    TotalAvailableBalance  string              `json:"totalAvailableBalance"`
    AccountMMRate          string              `json:"accountMMRate"`
    TotalPerpUPL           string              `json:"totalPerpUPL"`
    TotalWalletBalance     float64             `json:"totalWalletBalance"`
    TotalMaintenanceMargin string              `json:"totalMaintenanceMargin"`
    Coin                   []WalletBalanceCoin `json:"coin"`
}

type WalletBalanceCoin struct {
    AvailableToBorrow   string `json:"availableToBorrow"`
    AccruedInterest     string `json:"accruedInterest"`
    AvailableToWithdraw string `json:"availableToWithdraw"`
    TotalOrderIM        string `json:"totalOrderIM"`
    Equity              string `json:"equity"`
    TotalPositionMM     string `json:"totalPositionMM"`
    UsdValue            string `json:"usdValue"`
    UnrealisedPnl       string `json:"unrealisedPnl"`
    BorrowAmount        string `json:"borrowAmount"`
    TotalPositionIM     string `json:"totalPositionIM"`
    WalletBalance       string `json:"walletBalance"`
    CumRealisedPnl      string `json:"cumRealisedPnl"`
    Coin                string `json:"coin"`
}
