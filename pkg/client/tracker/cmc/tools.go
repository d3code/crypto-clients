package cmc

import (
    "encoding/json"
    "net/url"
)

func (c *Configuration) GetToolsPriceConversion(q url.Values) (*ToolsPriceConversion, error) {
    requestURL := "/v2/tools/price-conversion"

    resBody, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var cmcResponse ToolsPriceConversion

    unmarshalError := json.Unmarshal(resBody.Body, &cmcResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &cmcResponse, nil
}
