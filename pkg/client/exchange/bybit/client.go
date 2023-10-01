package bybit

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "github.com/d3code/crypto-clients/pkg/client"
    "github.com/d3code/zlog"
    "net/http"
    "net/url"
    "strconv"
    "time"
)

const (
    baseUrl             = "https://api.bybit.com"
    window              = "5000"
    headerContentType   = "Content-Type"
    headerApiKey        = "X-BAPI-API-KEY"
    headerSignature     = "X-BAPI-SIGN"
    headerTimestamp     = "X-BAPI-TIMESTAMP"
    headerSignType      = "X-BAPI-SIGN-TYPE"
    headerReceiveWindow = "X-BAPI-RECV-WINDOW"
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
    }
    return c
}

func (c *Configuration) Replay(replay bool, replayPath string) *Configuration {
    c.replay = replay
    c.replayPath = replayPath
    return c
}

func request(c *Configuration, method string, endpoint string, params url.Values, try int) (*client.Response, error) {
    now := time.Now()
    timeStamp := now.UnixNano() / 1000000

    hmac256 := hmac.New(sha256.New, []byte(c.apiSecret))
    hmac256.Write([]byte(strconv.FormatInt(timeStamp, 10) + c.apiKey + window + params.Encode()))
    signature := hex.EncodeToString(hmac256.Sum(nil))

    req, _ := http.NewRequest(method, baseUrl+endpoint, nil)
    req.URL.RawQuery = params.Encode()

    req.Header.Set(headerContentType, "application/json")
    req.Header.Set(headerApiKey, c.apiKey)
    req.Header.Set(headerSignature, signature)
    req.Header.Set(headerTimestamp, strconv.FormatInt(timeStamp, 10))
    req.Header.Set(headerSignType, "2")
    req.Header.Set(headerReceiveWindow, window)

    response, requestError := client.DoRequest(req, c.replay, c.replayPath, []string{headerTimestamp, headerApiKey, headerSignature}, []string{})
    if requestError != nil {
        return nil, requestError
    }

    if response.StatusCode == http.StatusTooManyRequests {
        if try > 3 {
            zlog.Log.Errorf("Too many requests\n%v\n%v", response.Headers, string(response.Body))
            return nil, fmt.Errorf("too many requests")
        }

        time.Sleep(60 * time.Second)
        return request(c, method, endpoint, params, try+1)
    }

    return response, nil
}

type ByBitResponse struct {
    RetCode    int    `json:"retCode"`
    RetMsg     string `json:"retMsg"`
    RetExtInfo struct {
    } `json:"retExtInfo"`
    Time int64 `json:"time"`
}
