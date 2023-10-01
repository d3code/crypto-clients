package coin_gecko

import (
    "github.com/d3code/crypto-clients/pkg/client"
    "net/http"
    "net/url"
)

const (
    cmcBaseUrl = "https://api.coingecko.com"
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

func doGetRequest(requestUrl string, q url.Values, c *Configuration) (*client.Response, error) {
    var u = cmcBaseUrl + requestUrl
    req, err := http.NewRequest(http.MethodGet, u, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Accepts", "application/json")
    req.URL.RawQuery = q.Encode()

    resBody, requestError := client.DoRequest(req, c.replay, c.replayPath, []string{}, []string{})
    if requestError != nil {
        return nil, requestError
    }

    return resBody, nil
}
