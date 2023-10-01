package cmc

import "time"

type Exchange struct {
    Id                  int       `json:"id"`
    Name                string    `json:"name"`
    Slug                string    `json:"slug"`
    IsActive            int       `json:"is_active"`
    IsListed            int       `json:"is_listed"`
    IsRedistributable   int       `json:"is_redistributable"`
    FirstHistoricalData time.Time `json:"first_historical_data"`
    LastHistoricalData  time.Time `json:"last_historical_data"`
}

type ExchangeMap struct {
    Data   []Exchange `json:"data"`
    Status Status     `json:"status"`
}

type Urls struct {
    Website []string      `json:"website"`
    Twitter []string      `json:"twitter"`
    Blog    []interface{} `json:"blog"`
    Chat    []string      `json:"chat"`
    Fee     []string      `json:"fee"`
}

type ExchangeInfo struct {
    Id                    int           `json:"id"`
    Name                  string        `json:"name"`
    Slug                  string        `json:"slug"`
    Logo                  string        `json:"logo"`
    Description           string        `json:"description"`
    DateLaunched          time.Time     `json:"date_launched"`
    Notice                interface{}   `json:"notice"`
    Countries             []interface{} `json:"countries"`
    Fiats                 []string      `json:"fiats"`
    Tags                  interface{}   `json:"tags"`
    Type                  string        `json:"type"`
    MakerFee              float64       `json:"maker_fee"`
    TakerFee              float64       `json:"taker_fee"`
    WeeklyVisits          int           `json:"weekly_visits"`
    SpotVolumeUsd         float64       `json:"spot_volume_usd"`
    SpotVolumeLastUpdated time.Time     `json:"spot_volume_last_updated"`
    PorStatus             int           `json:"porStatus"`
    PorAuditStatus        int           `json:"porAuditStatus"`
    Urls                  Urls          `json:"urls"`
}

type ExchangeInfoResponse struct {
    Data   map[string]ExchangeInfo `json:"data"`
    Status Status                  `json:"status"`
}

type ExchangeListingsLatest struct {
    Data []struct {
        Id             int       `json:"id"`
        Name           string    `json:"name"`
        Slug           string    `json:"slug"`
        NumMarketPairs int       `json:"num_market_pairs"`
        Fiats          []string  `json:"fiats"`
        TrafficScore   float64   `json:"traffic_score"`
        Rank           int       `json:"rank"`
        ExchangeScore  float64   `json:"exchange_score"`
        LiquidityScore float64   `json:"liquidity_score"`
        LastUpdated    time.Time `json:"last_updated"`
        Quote          map[string]struct {
            Volume24H              float64 `json:"volume_24h"`
            Volume24HAdjusted      float64 `json:"volume_24h_adjusted"`
            Volume7D               int64   `json:"volume_7d"`
            Volume30D              int64   `json:"volume_30d"`
            PercentChangeVolume24H float64 `json:"percent_change_volume_24h"`
            PercentChangeVolume7D  float64 `json:"percent_change_volume_7d"`
            PercentChangeVolume30D float64 `json:"percent_change_volume_30d"`
            EffectiveLiquidity24H  float64 `json:"effective_liquidity_24h"`
            DerivativeVolumeUsd    float64 `json:"derivative_volume_usd"`
            SpotVolumeUsd          float64 `json:"spot_volume_usd"`
        } `json:"quote"`
    } `json:"data"`
    Status Status `json:"status"`
}

type ExchangeMarketPairsLatest struct {
    Data struct {
        Id             int     `json:"id"`
        Name           string  `json:"name"`
        Slug           string  `json:"slug"`
        NumMarketPairs int     `json:"num_market_pairs"`
        Volume24H      float64 `json:"volume_24h"`
        MarketPairs    []struct {
            MarketId        int         `json:"market_id"`
            MarketPair      string      `json:"market_pair"`
            Category        string      `json:"category"`
            FeeType         string      `json:"fee_type"`
            OutlierDetected int         `json:"outlier_detected"`
            Exclusions      interface{} `json:"exclusions"`
            MarketPairBase  struct {
                CurrencyId     int    `json:"currency_id"`
                CurrencySymbol string `json:"currency_symbol"`
                ExchangeSymbol string `json:"exchange_symbol"`
                CurrencyType   string `json:"currency_type"`
            } `json:"market_pair_base"`
            MarketPairQuote struct {
                CurrencyId     int    `json:"currency_id"`
                CurrencySymbol string `json:"currency_symbol"`
                ExchangeSymbol string `json:"exchange_symbol"`
                CurrencyType   string `json:"currency_type"`
            } `json:"market_pair_quote"`
            Quote map[string]struct {
                Price            float64   `json:"price"`
                Volume24H        float64   `json:"volume_24h"`
                Volume24HBase    float64   `json:"volume_24h_base"`
                Volume24HQuote   float64   `json:"volume_24h_quote"`
                VolumePercentage float64   `json:"volume_percentage"`
                DepthNegativeTwo float64   `json:"depth_negative_two"`
                DepthPositiveTwo float64   `json:"depth_positive_two"`
                LastUpdated      time.Time `json:"last_updated"`
            } `json:"quote"`
        } `json:"market_pairs"`
    } `json:"data"`
    Status Status `json:"status"`
}

type ExchangeQuotesHistorical struct {
    Data struct {
        Id     int    `json:"id"`
        Name   string `json:"name"`
        Slug   string `json:"slug"`
        Quotes []struct {
            Timestamp time.Time `json:"timestamp"`
            Quote     map[string]struct {
                Volume24H int       `json:"volume_24h"`
                Timestamp time.Time `json:"timestamp"`
            } `json:"quote"`
            NumMarketPairs int `json:"num_market_pairs"`
        } `json:"quotes"`
    } `json:"data"`
    Status Status `json:"status"`
}

type ExchangeQuotesLatest struct {
    Data map[string]struct {
        Id             int       `json:"id"`
        Name           string    `json:"name"`
        Slug           string    `json:"slug"`
        NumCoins       int       `json:"num_coins"`
        NumMarketPairs int       `json:"num_market_pairs"`
        LastUpdated    time.Time `json:"last_updated"`
        TrafficScore   int       `json:"traffic_score"`
        Rank           int       `json:"rank"`
        ExchangeScore  float64   `json:"exchange_score"`
        LiquidityScore float64   `json:"liquidity_score"`
        Quote          map[string]struct {
            Volume24H              float64 `json:"volume_24h"`
            Volume24HAdjusted      float64 `json:"volume_24h_adjusted"`
            Volume7D               int64   `json:"volume_7d"`
            Volume30D              int64   `json:"volume_30d"`
            PercentChangeVolume24H float64 `json:"percent_change_volume_24h"`
            PercentChangeVolume7D  float64 `json:"percent_change_volume_7d"`
            PercentChangeVolume30D float64 `json:"percent_change_volume_30d"`
            EffectiveLiquidity24H  float64 `json:"effective_liquidity_24h"`
        } `json:"quote"`
    } `json:"data"`
    Status Status `json:"status"`
}
