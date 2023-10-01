package cmc

import "time"

type CryptocurrencyAirdrop struct {
    Data struct {
        Id          string `json:"id"`
        ProjectName string `json:"project_name"`
        Description string `json:"description"`
        Status      string `json:"status"`
        Coin        struct {
            Id     int    `json:"id"`
            Name   string `json:"name"`
            Slug   string `json:"slug"`
            Symbol string `json:"symbol"`
        } `json:"coin"`
        StartDate   time.Time `json:"start_date"`
        EndDate     time.Time `json:"end_date"`
        TotalPrize  int64     `json:"total_prize"`
        WinnerCount int       `json:"winner_count"`
        Link        string    `json:"link"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyAirdrops struct {
    Data struct {
        Data []struct {
            Id          string `json:"id"`
            ProjectName string `json:"project_name"`
            Description string `json:"description"`
            Status      string `json:"status"`
            Coin        struct {
                Id     int    `json:"id"`
                Name   string `json:"name"`
                Slug   string `json:"slug"`
                Symbol string `json:"symbol"`
            } `json:"coin"`
            StartDate   time.Time `json:"start_date"`
            EndDate     time.Time `json:"end_date"`
            TotalPrize  int       `json:"total_prize"`
            WinnerCount int       `json:"winner_count"`
            Link        string    `json:"link"`
        } `json:"data"`
        Status Status `json:"status"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyCategories struct {
    Status Status `json:"status"`
    Data   struct {
        Data []struct {
            Id              string  `json:"id"`
            Name            string  `json:"name"`
            Title           string  `json:"title"`
            Description     string  `json:"description"`
            NumTokens       int     `json:"num_tokens"`
            AvgPriceChange  float64 `json:"avg_price_change"`
            MarketCap       float64 `json:"market_cap"`
            MarketCapChange float64 `json:"market_cap_change"`
            Volume          float64 `json:"volume"`
            VolumeChange    float64 `json:"volume_change"`
            LastUpdated     int     `json:"last_updated"`
        } `json:"data"`
        Status Status `json:"status"`
    } `json:"data"`
}

type CryptocurrencyCategory struct {
    Data struct {
        Id              string  `json:"id"`
        Name            string  `json:"name"`
        Title           string  `json:"title"`
        Description     string  `json:"description"`
        NumTokens       int     `json:"num_tokens"`
        AvgPriceChange  float64 `json:"avg_price_change"`
        MarketCap       float64 `json:"market_cap"`
        MarketCapChange float64 `json:"market_cap_change"`
        Volume          float64 `json:"volume"`
        VolumeChange    float64 `json:"volume_change"`
        Coins           []struct {
            Id                int         `json:"id"`
            Name              string      `json:"name"`
            Symbol            string      `json:"symbol"`
            Slug              string      `json:"slug"`
            CmcRank           int         `json:"cmc_rank,omitempty"`
            NumMarketPairs    int         `json:"num_market_pairs"`
            CirculatingSupply int         `json:"circulating_supply"`
            TotalSupply       int         `json:"total_supply"`
            MaxSupply         int         `json:"max_supply"`
            LastUpdated       time.Time   `json:"last_updated"`
            DateAdded         time.Time   `json:"date_added"`
            Tags              []string    `json:"tags"`
            Platform          interface{} `json:"platform"`
            Quote             map[string]struct {
                Price            float64   `json:"price"`
                Volume24H        int64     `json:"volume_24h"`
                PercentChange1H  float64   `json:"percent_change_1h"`
                PercentChange24H float64   `json:"percent_change_24h"`
                PercentChange7D  float64   `json:"percent_change_7d"`
                MarketCap        int64     `json:"market_cap"`
                LastUpdated      time.Time `json:"last_updated"`
            } `json:"quote"`
        } `json:"coins"`
        LastUpdated int64 `json:"last_updated"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyInfo struct {
    Data map[string]struct {
        Id          int    `json:"id"`
        Name        string `json:"name"`
        Symbol      string `json:"symbol"`
        Slug        string `json:"slug"`
        Description string `json:"description"`

        Status string `json:"status"`
        Hidden int    `json:"is_hidden"`

        Logo     string `json:"logo"`
        Category string `json:"category"`

        DateAdded    time.Time `json:"date_added"`
        DateLaunched time.Time `json:"date_launched"`

        Subreddit       string `json:"subreddit"`
        TwitterUsername string `json:"twitter_username"`
        Notice          string `json:"notice"`

        Tags      []string `json:"tags"`
        TagNames  []string `json:"tag-names"`
        TagGroups []string `json:"tag-groups"`

        Platform struct {
            Id           string `json:"id"`
            Name         string `json:"name"`
            Slug         string `json:"slug"`
            Symbol       string `json:"symbol"`
            TokenAddress string `json:"token_address"`
        } `json:"platform"`

        Urls struct {
            Website      []string `json:"website"`
            TechnicalDoc []string `json:"technical_doc"`
            Twitter      []string `json:"twitter"`
            Reddit       []string `json:"reddit"`
            MessageBoard []string `json:"message_board"`
            Announcement []string `json:"announcement"`
            Chat         []string `json:"chat"`
            Explorer     []string `json:"explorer"`
            SourceCode   []string `json:"source_code"`
        } `json:"urls"`

        ContractAddress []struct {
            ContractAddress string `json:"contract_address"`
            Platform        struct {
                Name string `json:"name"`
                Coin struct {
                    Id     string `json:"id"`
                    Name   string `json:"name"`
                    Symbol string `json:"symbol"`
                    Slug   string `json:"slug"`
                } `json:"coin"`
            } `json:"platform"`
        } `json:"contract_address"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyListingsHistorical struct {
    Data []struct {
        Id                int         `json:"id"`
        Name              string      `json:"name"`
        Symbol            string      `json:"symbol"`
        Slug              string      `json:"slug"`
        CmcRank           int         `json:"cmc_rank,omitempty"`
        NumMarketPairs    int         `json:"num_market_pairs"`
        CirculatingSupply int         `json:"circulating_supply"`
        TotalSupply       int         `json:"total_supply"`
        MaxSupply         int         `json:"max_supply"`
        LastUpdated       time.Time   `json:"last_updated"`
        DateAdded         time.Time   `json:"date_added"`
        Tags              []string    `json:"tags"`
        Platform          interface{} `json:"platform"`
        Quote             map[string]struct {
            Price            float64   `json:"price"`
            Volume24H        int64     `json:"volume_24h"`
            PercentChange1H  float64   `json:"percent_change_1h"`
            PercentChange24H float64   `json:"percent_change_24h"`
            PercentChange7D  float64   `json:"percent_change_7d"`
            MarketCap        int64     `json:"market_cap"`
            LastUpdated      time.Time `json:"last_updated"`
        } `json:"quote"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyListingsLatest struct {
    Data []struct {
        Id                            int         `json:"id"`
        Name                          string      `json:"name"`
        Symbol                        string      `json:"symbol"`
        Slug                          string      `json:"slug"`
        CmcRank                       int         `json:"cmc_rank,omitempty"`
        NumMarketPairs                int         `json:"num_market_pairs"`
        CirculatingSupply             int         `json:"circulating_supply"`
        TotalSupply                   int         `json:"total_supply"`
        MaxSupply                     int         `json:"max_supply"`
        LastUpdated                   time.Time   `json:"last_updated"`
        DateAdded                     time.Time   `json:"date_added"`
        Tags                          []string    `json:"tags"`
        Platform                      interface{} `json:"platform"`
        SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
        SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
        Quote                         map[string]struct {
            Price                 float64   `json:"price"`
            Volume24H             int64     `json:"volume_24h"`
            VolumeChange24H       float64   `json:"volume_change_24h"`
            PercentChange1H       float64   `json:"percent_change_1h"`
            PercentChange24H      float64   `json:"percent_change_24h"`
            PercentChange7D       float64   `json:"percent_change_7d"`
            MarketCap             float64   `json:"market_cap"`
            MarketCapDominance    int       `json:"market_cap_dominance"`
            FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
            LastUpdated           time.Time `json:"last_updated"`
        } `json:"quote"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyListingsNew struct {
    Data []struct {
        Id                int         `json:"id"`
        Name              string      `json:"name"`
        Symbol            string      `json:"symbol"`
        Slug              string      `json:"slug"`
        CmcRank           int         `json:"cmc_rank,omitempty"`
        NumMarketPairs    int         `json:"num_market_pairs"`
        CirculatingSupply int         `json:"circulating_supply"`
        TotalSupply       int         `json:"total_supply"`
        MaxSupply         int         `json:"max_supply"`
        LastUpdated       time.Time   `json:"last_updated"`
        DateAdded         time.Time   `json:"date_added"`
        Tags              []string    `json:"tags"`
        Platform          interface{} `json:"platform"`
        Quote             map[string]struct {
            Price                 float64   `json:"price"`
            Volume24H             int64     `json:"volume_24h"`
            VolumeChange24H       float64   `json:"volume_change_24h"`
            PercentChange1H       float64   `json:"percent_change_1h"`
            PercentChange24H      float64   `json:"percent_change_24h"`
            PercentChange7D       float64   `json:"percent_change_7d"`
            MarketCap             float64   `json:"market_cap"`
            MarketCapDominance    int       `json:"market_cap_dominance"`
            FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
            LastUpdated           time.Time `json:"last_updated"`
        } `json:"quote"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyTrendingGainersLosers struct {
    Data struct {
        Data []struct {
            Id                int         `json:"id"`
            Name              string      `json:"name"`
            Symbol            string      `json:"symbol"`
            Slug              string      `json:"slug"`
            CmcRank           int         `json:"cmc_rank"`
            NumMarketPairs    int         `json:"num_market_pairs"`
            CirculatingSupply int         `json:"circulating_supply"`
            TotalSupply       int         `json:"total_supply"`
            MaxSupply         int         `json:"max_supply"`
            LastUpdated       time.Time   `json:"last_updated"`
            DateAdded         time.Time   `json:"date_added"`
            Tags              []string    `json:"tags"`
            Platform          interface{} `json:"platform"`
            Quote             map[string]struct {
                Price            float64   `json:"price"`
                Volume24H        int       `json:"volume_24h"`
                PercentChange1H  float64   `json:"percent_change_1h"`
                PercentChange24H float64   `json:"percent_change_24h"`
                PercentChange7D  float64   `json:"percent_change_7d"`
                MarketCap        int       `json:"market_cap"`
                LastUpdated      time.Time `json:"last_updated"`
            } `json:"quote"`
        } `json:"data"`
        Status Status `json:"status"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyTrendingLatest struct {
    Data struct {
        Data []struct {
            Id                            int         `json:"id"`
            Name                          string      `json:"name"`
            Symbol                        string      `json:"symbol"`
            Slug                          string      `json:"slug"`
            CmcRank                       int         `json:"cmc_rank"`
            IsActive                      bool        `json:"is_active"`
            IsFiat                        interface{} `json:"is_fiat"`
            SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
            SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
            NumMarketPairs                int         `json:"num_market_pairs"`
            CirculatingSupply             int         `json:"circulating_supply"`
            TotalSupply                   int         `json:"total_supply"`
            MaxSupply                     int         `json:"max_supply"`
            LastUpdated                   time.Time   `json:"last_updated"`
            DateAdded                     time.Time   `json:"date_added"`
            Tags                          []string    `json:"tags"`
            Platform                      interface{} `json:"platform"`
            Quote                         map[string]struct {
                Price            float64   `json:"price"`
                Volume24H        int       `json:"volume_24h"`
                PercentChange1H  float64   `json:"percent_change_1h"`
                PercentChange24H float64   `json:"percent_change_24h"`
                PercentChange7D  float64   `json:"percent_change_7d"`
                MarketCap        int       `json:"market_cap"`
                LastUpdated      time.Time `json:"last_updated"`
            } `json:"quote"`
        } `json:"data"`
        Status Status `json:"status"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyTrendingMostVisited struct {
    Data struct {
        Data []struct {
            Id                int         `json:"id"`
            Name              string      `json:"name"`
            Symbol            string      `json:"symbol"`
            Slug              string      `json:"slug"`
            CmcRank           int         `json:"cmc_rank"`
            NumMarketPairs    int         `json:"num_market_pairs"`
            CirculatingSupply int         `json:"circulating_supply"`
            TotalSupply       int         `json:"total_supply"`
            MaxSupply         int         `json:"max_supply"`
            LastUpdated       time.Time   `json:"last_updated"`
            DateAdded         time.Time   `json:"date_added"`
            Tags              []string    `json:"tags"`
            Platform          interface{} `json:"platform"`
            Quote             map[string]struct {
                Price            float64   `json:"price"`
                Volume24H        int       `json:"volume_24h"`
                PercentChange1H  float64   `json:"percent_change_1h"`
                PercentChange24H float64   `json:"percent_change_24h"`
                PercentChange7D  float64   `json:"percent_change_7d"`
                MarketCap        int       `json:"market_cap"`
                LastUpdated      time.Time `json:"last_updated"`
            } `json:"quote"`
        } `json:"data"`
        Status Status `json:"status"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyMarketPairsLatest struct {
    Data struct {
        Id             int    `json:"id"`
        Name           string `json:"name"`
        Symbol         string `json:"symbol"`
        NumMarketPairs int    `json:"num_market_pairs"`
        MarketPairs    []struct {
            Exchange struct {
                Id   int    `json:"id"`
                Name string `json:"name"`
                Slug string `json:"slug"`
            } `json:"exchange"`
            MarketId       int    `json:"market_id"`
            MarketPair     string `json:"market_pair"`
            Category       string `json:"category"`
            FeeType        string `json:"fee_type"`
            MarketPairBase struct {
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
                Price          float64   `json:"price"`
                Volume24H      float64   `json:"volume_24h"`
                Volume24HBase  float64   `json:"volume_24h_base"`
                Volume24HQuote float64   `json:"volume_24h_quote"`
                LastUpdated    time.Time `json:"last_updated"`
            } `json:"quote"`
        } `json:"market_pairs"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyOHLCVHistorical struct {
    Data struct {
        Id     int    `json:"id"`
        Name   string `json:"name"`
        Symbol string `json:"symbol"`
        Quotes []struct {
            TimeOpen  time.Time `json:"time_open"`
            TimeClose time.Time `json:"time_close"`
            TimeHigh  time.Time `json:"time_high"`
            TimeLow   time.Time `json:"time_low"`
            Quote     map[string]struct {
                Open      float64   `json:"open"`
                High      float64   `json:"high"`
                Low       float64   `json:"low"`
                Close     float64   `json:"close"`
                Volume    float64   `json:"volume"`
                MarketCap float64   `json:"market_cap"`
                Timestamp time.Time `json:"timestamp"`
            } `json:"quote"`
        } `json:"quotes"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyOHLCVLatest struct {
    Data map[string]struct {
        Id          int         `json:"id"`
        Name        string      `json:"name"`
        Symbol      string      `json:"symbol"`
        LastUpdated time.Time   `json:"last_updated"`
        TimeOpen    time.Time   `json:"time_open"`
        TimeClose   interface{} `json:"time_close"`
        TimeHigh    time.Time   `json:"time_high"`
        TimeLow     time.Time   `json:"time_low"`
        Quote       map[string]struct {
            Open        float64   `json:"open"`
            High        float64   `json:"high"`
            Low         float64   `json:"low"`
            Close       float64   `json:"close"`
            Volume      int64     `json:"volume"`
            LastUpdated time.Time `json:"last_updated"`
        } `json:"quote"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyPricePerformanceStats struct {
    Data struct {
        Field1 struct {
            Id          int       `json:"id"`
            Name        string    `json:"name"`
            Symbol      string    `json:"symbol"`
            Slug        string    `json:"slug"`
            LastUpdated time.Time `json:"last_updated"`
            Periods     map[string]struct {
                OpenTimestamp  time.Time `json:"open_timestamp"`
                HighTimestamp  time.Time `json:"high_timestamp"`
                LowTimestamp   time.Time `json:"low_timestamp"`
                CloseTimestamp time.Time `json:"close_timestamp"`
                Quote          map[string]struct {
                    Open           float64   `json:"open"`
                    OpenTimestamp  time.Time `json:"open_timestamp"`
                    High           float64   `json:"high"`
                    HighTimestamp  time.Time `json:"high_timestamp"`
                    Low            float64   `json:"low"`
                    LowTimestamp   time.Time `json:"low_timestamp"`
                    Close          float64   `json:"close"`
                    CloseTimestamp time.Time `json:"close_timestamp"`
                    PercentChange  float64   `json:"percent_change"`
                    PriceChange    float64   `json:"price_change"`
                } `json:"quote"`
            } `json:"all_time"`
        } `json:"periods"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyQuotesHistorical struct {
    Data struct {
        Id       int    `json:"id"`
        Name     string `json:"name"`
        Symbol   string `json:"symbol"`
        IsActive int    `json:"is_active"`
        IsFiat   int    `json:"is_fiat"`
        Quotes   []struct {
            Timestamp time.Time `json:"timestamp"`
            Quote     map[string]struct {
                Price             float64   `json:"price"`
                Volume24H         int64     `json:"volume_24h"`
                MarketCap         float64   `json:"market_cap"`
                CirculatingSupply int64     `json:"circulating_supply"`
                TotalSupply       int64     `json:"total_supply"`
                Timestamp         time.Time `json:"timestamp"`
            } `json:"quote"`
        } `json:"quotes"`
    } `json:"data"`
    Status Status `json:"status"`
}

type CryptocurrencyQuotesLatest struct {
    Data map[string]struct {
        Id                            int         `json:"id"`
        Name                          string      `json:"name"`
        Symbol                        string      `json:"symbol"`
        Slug                          string      `json:"slug"`
        IsActive                      int         `json:"is_active"`
        IsFiat                        int         `json:"is_fiat"`
        CirculatingSupply             int         `json:"circulating_supply"`
        TotalSupply                   int         `json:"total_supply"`
        MaxSupply                     int         `json:"max_supply"`
        DateAdded                     time.Time   `json:"date_added"`
        NumMarketPairs                int         `json:"num_market_pairs"`
        CmcRank                       int         `json:"cmc_rank"`
        LastUpdated                   time.Time   `json:"last_updated"`
        Tags                          []string    `json:"tags"`
        Platform                      interface{} `json:"platform"`
        SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
        SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
        Quote                         map[string]struct {
            Price                 float64   `json:"price"`
            Volume24H             float64   `json:"volume_24h"`
            VolumeChange24H       float64   `json:"volume_change_24h"`
            PercentChange1H       float64   `json:"percent_change_1h"`
            PercentChange24H      float64   `json:"percent_change_24h"`
            PercentChange7D       float64   `json:"percent_change_7d"`
            PercentChange30D      float64   `json:"percent_change_30d"`
            MarketCap             float64   `json:"market_cap"`
            MarketCapDominance    int       `json:"market_cap_dominance"`
            FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
            LastUpdated           time.Time `json:"last_updated"`
        } `json:"quote"`
    } `json:"data"`
    Status Status `json:"status"`
}
