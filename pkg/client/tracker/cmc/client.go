package cmc

import (
    "github.com/d3code/crypto-clients/pkg/client"
    "net/http"
    "net/url"
    "time"
)

const (
    cmcBaseUrl      = "https://pro-api.coinmarketcap.com"
    cmcApiKeyHeader = "X-CMC_PRO_API_KEY"
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

func (c *Configuration) Replay(replay bool, replayPath string) *Configuration {
    c.replay = replay
    c.replayPath = replayPath
    return c
}

func request(requestUrl string, q url.Values, c *Configuration) (*client.Response, error) {
    req, err := http.NewRequest(http.MethodGet, cmcBaseUrl+requestUrl, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set(cmcApiKeyHeader, c.apiKey)
    req.Header.Set("Accepts", "application/json")
    req.URL.RawQuery = q.Encode()

    return client.DoRequest(req, c.replay, c.replayPath, []string{cmcApiKeyHeader}, []string{})
}

type Status struct {
    Timestamp    time.Time `json:"timestamp"`
    ErrorCode    int       `json:"error_code"`
    ErrorMessage string    `json:"error_message"`
    Elapsed      int       `json:"elapsed"`
    CreditCount  int       `json:"credit_count"`
    Notice       string    `json:"notice"`
}
