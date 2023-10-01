package binance

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "github.com/d3code/crypto-clients/pkg/client"
    "net/http"
    "net/url"
    "strconv"
    "time"
)

const (
    cmcBaseUrl     = "https://api.binance.com"
    headerApiKey   = "X-MBX-APIKEY"
    paramTimestamp = "timestamp"
    paramSignature = "signature"
)

type Configuration struct {
    apiKey     string
    apiSecret  string
    replay     bool
    replayPath string
}

func Client(apiKey string, apiSecret string) *Configuration {
    c := &Configuration{
        apiKey:    apiKey,
        apiSecret: apiSecret,
        replay:    false,
    }

    return c
}

func (c *Configuration) Replay(replay bool, replayPath string) *Configuration {
    c.replay = replay
    c.replayPath = replayPath
    return c
}

func request(httpMethod string, requestUrl string, values url.Values, c *Configuration) (*client.Response, error) {
    baseUrl := cmcBaseUrl + requestUrl

    httpRequest, err := http.NewRequest(httpMethod, baseUrl, nil)
    if err != nil {
        return nil, err
    }

    httpRequest.Header.Set("Accepts", "application/json")
    httpRequest.Header.Set(headerApiKey, c.apiKey)

    values.Add(paramTimestamp, strconv.FormatInt(time.Now().UnixMilli(), 10))

    shaHmac := hmac.New(sha256.New, []byte(c.apiSecret))
    shaHmac.Write([]byte(values.Encode()))
    sha := hex.EncodeToString(shaHmac.Sum(nil))

    values.Add(paramSignature, sha)

    httpRequest.URL.RawQuery = values.Encode()

    response, err := client.DoRequest(httpRequest, c.replay, c.replayPath, []string{headerApiKey}, []string{paramSignature, paramTimestamp})
    if err != nil {
        return nil, err
    }

    return response, nil
}
