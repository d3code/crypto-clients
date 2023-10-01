package bybit

import (
    "encoding/json"
    "github.com/d3code/zlog"
    "net/http"
    "net/url"
    "strconv"
    "time"
)

// GetClosedPnL returns the closed PnL for a given symbol.
// https://bybit-exchange.github.io/docs/v5/position/close-pnl
//
// symbol: linear, inverse
func (c *Configuration) GetClosedPnL(category string, from time.Time) ([]ClosedPnL, error) {
    var closedPnLS []ClosedPnL
    lv5, err := c.getClosedPnL(category, from, "")
    if err != nil {
        return nil, err
    }

    closedPnLS = append(closedPnLS, lv5.Result.List...)
    if lv5.Result.NextPageCursor != "" {
        for {
            lv5, err = c.getClosedPnL(category, from, lv5.Result.NextPageCursor)
            if err != nil {
                return nil, err
            }
            closedPnLS = append(closedPnLS, lv5.Result.List...)
            if lv5.Result.NextPageCursor == "" {
                break
            }
        }
    }

    return closedPnLS, err
}

func (c *Configuration) getClosedPnL(category string, from time.Time, cursor string) (*ClosedPnLResponse, error) {
    requestURL := "/v5/position/closed-pnl"

    values := url.Values{}
    values.Add("category", category)
    values.Add("limit", "100")
    values.Add("startTime", strconv.FormatInt(from.UnixMilli(), 10))

    if cursor != "" {
        values.Add("cursor", cursor)
    }

    response, err := request(c, http.MethodGet, requestURL, values, 0)
    if err != nil {
        return nil, err
    }

    var responseObject ClosedPnLResponse

    unmarshalError := json.Unmarshal(response.Body, &responseObject)
    if unmarshalError != nil {
        zlog.Log.Errorf("unmarshalError: %v", unmarshalError)
        return nil, unmarshalError
    }

    zlog.Log.Debugf("GetClosedPnLV5: %+v", responseObject.Result.List)

    return &responseObject, nil
}

type ClosedPnLResponse struct {
    ByBitResponse
    Result struct {
        NextPageCursor string      `json:"nextPageCursor"`
        Category       string      `json:"category"`
        List           []ClosedPnL `json:"list"`
    } `json:"result"`
}

type ClosedPnL struct {
    Symbol        string `json:"symbol"`
    OrderType     string `json:"orderType"`
    Leverage      string `json:"leverage"`
    UpdatedTime   string `json:"updatedTime"`
    Side          string `json:"side"`
    OrderId       string `json:"orderId"`
    ClosedPnl     string `json:"closedPnl"`
    AvgEntryPrice string `json:"avgEntryPrice"`
    Qty           string `json:"qty"`
    CumEntryValue string `json:"cumEntryValue"`
    CreatedTime   string `json:"createdTime"`
    OrderPrice    string `json:"orderPrice"`
    ClosedSize    string `json:"closedSize"`
    AvgExitPrice  string `json:"avgExitPrice"`
    ExecType      string `json:"execType"`
    FillCount     string `json:"fillCount"`
    CumExitValue  string `json:"cumExitValue"`
}
