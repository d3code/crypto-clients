package client

import (
    "github.com/d3code/zlog"
    "io"
    "net/http"
    "strings"
    "time"
)

type Response struct {
    StatusCode    int               `json:"status_code"`
    Headers       map[string]string `json:"headers"`
    ContentLength int64             `json:"content_length"`
    Body          []byte            `json:"body"`
    Stored        bool              `json:"stored"`
}

func getRequestEntry(req *http.Request, removeQueryParam []string, removeHeader []string) RequestEntry {
    host := strings.Replace(req.URL.Host, ".", "_", -1)
    path := req.URL.EscapedPath()
    params := getParams(req, removeQueryParam)
    headers := getHeaders(req, removeHeader)

    now := time.Now()
    requestEntry := RequestEntry{
        Host:      host,
        Path:      path,
        Params:    params,
        Headers:   headers,
        Method:    req.Method,
        Timestamp: now,
        File:      now.Format("2006-01-02_15:04:05_.999999999"),
    }

    return requestEntry
}

func getResponse(req *http.Request) (*Response, error) {
    httpClient := http.Client{
        Timeout: 30 * time.Second,
    }

    response, requestError := httpClient.Do(req)
    if requestError != nil {
        zlog.Log.Errorf("Error while executing request: %s", requestError.Error())
        return nil, requestError
    }

    responseBody, responseError := io.ReadAll(response.Body)
    if responseError != nil {
        zlog.Log.Errorf("Error while reading response: %s", responseError.Error())
        return nil, responseError
    }

    var headers = make(map[string]string)
    for key, value := range response.Header {
        headers[key] = strings.Join(value, ";")
    }

    responseInfo := Response{
        StatusCode:    response.StatusCode,
        Headers:       headers,
        ContentLength: response.ContentLength,
        Body:          responseBody,
        Stored:        false,
    }

    return &responseInfo, nil
}
