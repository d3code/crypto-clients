package cmc

import (
    "encoding/json"
    "net/url"
)

func (c *Configuration) GetKeyInfo(q url.Values) (*KeyInfo, error) {
    requestURL := "/v1/key/info"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse KeyInfo

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}
