package binance

import (
    "encoding/json"
    "github.com/d3code/zlog"
    "net/http"
    "net/url"
    "strconv"
    "time"
)

func (c *Configuration) GetDeposits() (*[]Deposit, error) {
    var responseObject []Deposit

    now := time.Now()

    for times := 1; times <= 8; times++ {
        deposits, err := c.getDepositsPastMonth(times, now)
        if err != nil {
            return nil, err
        }

        if len(*deposits) == 0 {
            continue
        }

        responseObject = append(responseObject, *deposits...)
    }

    return &responseObject, nil
}

func (c *Configuration) getDepositsPastMonth(monthBack int, dt time.Time) (*[]Deposit, error) {
    requestURL := "/sapi/v1/capital/deposit/hisrec"

    from := dt.Add(time.Duration(monthBack) * time.Hour * 24 * 90 * -1)
    to := from.Add(time.Hour * 24 * 90)

    zlog.Log.Infof("Getting deposits from %v", from)

    values := url.Values{}
    values.Add("startTime", strconv.FormatInt(from.Unix()*1000, 10))
    values.Add("endTime", strconv.FormatInt(to.Unix()*1000, 10))

    response, err := request(http.MethodGet, requestURL, values, c)
    if err != nil {
        zlog.Log.Errorf("Error getting deposit: %s", err.Error())
        return nil, err
    }

    var responseObject []Deposit

    unmarshalError := json.Unmarshal(response.Body, &responseObject)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &responseObject, nil
}

type Deposit struct {
    Id            string `json:"id"`
    Amount        string `json:"amount"`
    Coin          string `json:"coin"`
    Network       string `json:"network"`
    Status        int    `json:"status"`
    Address       string `json:"address"`
    AddressTag    string `json:"addressTag"`
    TxId          string `json:"txId"`
    InsertTime    int64  `json:"insertTime"`
    TransferType  int    `json:"transferType"`
    ConfirmTimes  string `json:"confirmTimes"`
    UnlockConfirm int    `json:"unlockConfirm"`
    WalletType    int    `json:"walletType"`
}
