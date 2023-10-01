package bsc

import (
    "encoding/json"
    "net/url"
)

func (c *Configuration) GetAccountTokenTx(values url.Values) (*AccountTokenTx, error) {
    values.Add("module", "account")
    values.Add("action", "tokentx")
    values.Add("sort", "asc")

    resBody, err := doGetRequest(values, c)
    if err != nil {
        return nil, err
    }

    var response AccountTokenTx

    unmarshalError := json.Unmarshal(resBody, &response)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &response, nil
}

func (c *Configuration) GetAccountTxList(values url.Values) (*AccountTxList, error) {
    values.Add("module", "account")
    values.Add("action", "txlist")
    values.Add("sort", "asc")

    resBody, err := doGetRequest(values, c)
    if err != nil {
        return nil, err
    }

    var response AccountTxList

    unmarshalError := json.Unmarshal(resBody, &response)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &response, nil
}
