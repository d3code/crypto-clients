package asx

import (
    "github.com/d3code/crypto-clients/pkg/client"
    "net/http"
    "net/url"
)

const (
    baseUrl = "https://asx.api.markitdigital.com"
)

type Configuration struct {
    replay     bool
    replayPath string
}

func Client() *Configuration {
    c := &Configuration{
        replay: false,
    }

    return c
}

func (c *Configuration) Replay(replay bool, replayPath string) *Configuration {
    c.replay = replay
    c.replayPath = replayPath
    return c
}

func request(httpMethod string, path string, values url.Values, c *Configuration) (*client.Response, error) {
    requestUrl := baseUrl + path

    httpRequest, err := http.NewRequest(httpMethod, requestUrl, nil)
    if err != nil {
        return nil, err
    }

    httpRequest.Header.Set("Accepts", "application/json")
    httpRequest.URL.RawQuery = values.Encode()

    response, err := client.DoRequest(httpRequest, c.replay, c.replayPath, []string{}, []string{})
    if err != nil {
        return nil, err
    }

    return response, nil
}
