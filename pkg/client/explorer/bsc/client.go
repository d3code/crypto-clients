package bsc

import (
    "io"
    "net/http"
    "net/url"
    "time"
)

const baseUrl = "https://api.bscscan.com/api"

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

func doGetRequest(q url.Values, c *Configuration) ([]byte, error) {
    req, err := http.NewRequest(http.MethodGet, baseUrl, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Accepts", "application/json")

    q.Add("apikey", c.apiKey)
    req.URL.RawQuery = q.Encode()

    httpClient := http.Client{
        Timeout: 30 * time.Second,
    }

    res, requestError := httpClient.Do(req)

    if requestError != nil {
        return nil, requestError
    }

    resBody, responseError := io.ReadAll(res.Body)
    if responseError != nil {
        return nil, responseError
    }

    return resBody, nil
}
