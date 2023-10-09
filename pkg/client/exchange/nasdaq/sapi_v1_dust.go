package nasdaq

import (
    "encoding/json"
    "net/http"
    "net/url"
)

func (c *Configuration) GetDustLog() (*DustLog, error) {
    requestURL := "/sapi/v1/asset/dribblet"

    values := url.Values{}
    response, err := request(http.MethodGet, requestURL, values, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse DustLog
    unmarshalError := json.Unmarshal(response.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}

type DustLog struct {
    Total              int       `json:"total"`
    UserAssetDribblets []Driblet `json:"userAssetDribblets"`
}

type Driblet struct {
    OperateTime              int64            `json:"operateTime"`
    TotalServiceChargeAmount string           `json:"totalServiceChargeAmount"`
    TotalTransferedAmount    string           `json:"totalTransferedAmount"`
    TransId                  int64            `json:"transId"`
    UserAssetDribbletDetails []DribletDetails `json:"userAssetDribbletDetails"`
    ClientId                 string           `json:"clientId,omitempty"`
}

type DribletDetails struct {
    FromAsset           string `json:"fromAsset"`
    Amount              string `json:"amount"`
    TransferedAmount    string `json:"transferedAmount"`
    ServiceChargeAmount string `json:"serviceChargeAmount"`
    OperateTime         int64  `json:"operateTime"`
    TransId             int64  `json:"transId"`
}
