package etherscan

import (
    "github.com/d3code/crypto-clients/pkg/client"
    "net/http"
    "net/url"
)

const (
    baseUrl     = "https://api.etherscan.io/api"
    apiKeyParam = "apikey"
)

type Configuration struct {
    apiKey     string
    replay     bool
    replayPath string
}

func Client(apiKey string) *Configuration {
    c := &Configuration{
        apiKey: apiKey,
        replay: false,
    }
    return c
}

func doGetRequest(q url.Values, c *Configuration) (*client.Response, error) {
    req, err := http.NewRequest(http.MethodGet, baseUrl, nil)
    if err != nil {
        return nil, err
    }

    q.Add(apiKeyParam, c.apiKey)

    req.Header.Set("Accepts", "application/json")
    req.URL.RawQuery = q.Encode()

    response, requestError := client.DoRequest(req, c.replay, c.replayPath, []string{}, []string{apiKeyParam})
    if requestError != nil {
        return nil, requestError
    }

    return response, nil
}
