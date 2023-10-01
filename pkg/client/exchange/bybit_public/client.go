package bybit_public

import (
    "github.com/d3code/crypto-clients/pkg/client"
    "net/http"
    "net/url"
)

const (
    baseUrl = "https://api.bybit.com"
)

type Configuration struct {
    replay     bool
    replayPath string
}

func Client() *Configuration {
    c := &Configuration{}

    return c
}

func (c *Configuration) Replay(replay bool, replayPath string) *Configuration {
    c.replay = replay
    c.replayPath = replayPath
    return c
}

func request(httpMethod string, requestUrl string, values url.Values, c *Configuration) (*client.Response, error) {
    url := baseUrl + requestUrl

    req, err := http.NewRequest(httpMethod, url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Accepts", "application/json")

    req.URL.RawQuery = values.Encode()

    response, requestError := client.DoRequest(req, c.replay, c.replayPath, []string{}, []string{})
    if requestError != nil {
        return nil, requestError
    }

    return response, nil
}
