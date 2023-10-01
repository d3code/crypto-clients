package cmc

import "time"

type GlobalMetricsQuotesHistorical struct {
    Data struct {
        Quotes []struct {
            Timestamp              time.Time `json:"timestamp"`
            BtcDominance           float64   `json:"btc_dominance"`
            ActiveCryptocurrencies int       `json:"active_cryptocurrencies"`
            ActiveExchanges        int       `json:"active_exchanges"`
            ActiveMarketPairs      int       `json:"active_market_pairs"`
            Quote                  map[string]struct {
                TotalMarketCap           float64   `json:"total_market_cap"`
                TotalVolume24H           float64   `json:"total_volume_24h"`
                TotalVolume24HReported   int64     `json:"total_volume_24h_reported"`
                AltcoinMarketCap         int64     `json:"altcoin_market_cap"`
                AltcoinVolume24H         int64     `json:"altcoin_volume_24h"`
                AltcoinVolume24HReported int64     `json:"altcoin_volume_24h_reported"`
                Timestamp                time.Time `json:"timestamp"`
            } `json:"quote"`
        } `json:"quotes"`
    } `json:"data"`
    Status Status `json:"status"`
}

type GlobalMetricsQuotesLatest struct {
    Data struct {
        ActiveCryptocurrencies          int     `json:"active_cryptocurrencies"`
        TotalCryptocurrencies           int     `json:"total_cryptocurrencies"`
        ActiveMarketPairs               int     `json:"active_market_pairs"`
        ActiveExchanges                 int     `json:"active_exchanges"`
        TotalExchanges                  int     `json:"total_exchanges"`
        EthDominance                    float64 `json:"eth_dominance"`
        BtcDominance                    float64 `json:"btc_dominance"`
        EthDominanceYesterday           float64 `json:"eth_dominance_yesterday"`
        BtcDominanceYesterday           float64 `json:"btc_dominance_yesterday"`
        EthDominance24HPercentageChange float64 `json:"eth_dominance_24h_percentage_change"`
        BtcDominance24HPercentageChange float64 `json:"btc_dominance_24h_percentage_change"`
        DefiVolume24H                   float64 `json:"defi_volume_24h"`
        DefiVolume24HReported           float64 `json:"defi_volume_24h_reported"`
        DefiMarketCap                   float64 `json:"defi_market_cap"`
        Defi24HPercentageChange         float64 `json:"defi_24h_percentage_change"`
        StablecoinVolume24H             float64 `json:"stablecoin_volume_24h"`
        StablecoinVolume24HReported     float64 `json:"stablecoin_volume_24h_reported"`
        StablecoinMarketCap             float64 `json:"stablecoin_market_cap"`
        Stablecoin24HPercentageChange   float64 `json:"stablecoin_24h_percentage_change"`
        DerivativesVolume24H            float64 `json:"derivatives_volume_24h"`
        DerivativesVolume24HReported    float64 `json:"derivatives_volume_24h_reported"`
        Derivatives24HPercentageChange  float64 `json:"derivatives_24h_percentage_change"`
        Quote                           map[string]struct {
            TotalMarketCap                          float64   `json:"total_market_cap"`
            TotalVolume24H                          float64   `json:"total_volume_24h"`
            TotalVolume24HReported                  float64   `json:"total_volume_24h_reported"`
            AltcoinVolume24H                        float64   `json:"altcoin_volume_24h"`
            AltcoinVolume24HReported                float64   `json:"altcoin_volume_24h_reported"`
            AltcoinMarketCap                        float64   `json:"altcoin_market_cap"`
            DefiVolume24H                           float64   `json:"defi_volume_24h"`
            DefiVolume24HReported                   float64   `json:"defi_volume_24h_reported"`
            Defi24HPercentageChange                 float64   `json:"defi_24h_percentage_change"`
            DefiMarketCap                           float64   `json:"defi_market_cap"`
            StablecoinVolume24H                     float64   `json:"stablecoin_volume_24h"`
            StablecoinVolume24HReported             float64   `json:"stablecoin_volume_24h_reported"`
            Stablecoin24HPercentageChange           float64   `json:"stablecoin_24h_percentage_change"`
            StablecoinMarketCap                     float64   `json:"stablecoin_market_cap"`
            DerivativesVolume24H                    float64   `json:"derivatives_volume_24h"`
            DerivativesVolume24HReported            float64   `json:"derivatives_volume_24h_reported"`
            Derivatives24HPercentageChange          float64   `json:"derivatives_24h_percentage_change"`
            LastUpdated                             time.Time `json:"last_updated"`
            TotalMarketCapYesterday                 float64   `json:"total_market_cap_yesterday"`
            TotalVolume24HYesterday                 float64   `json:"total_volume_24h_yesterday"`
            TotalMarketCapYesterdayPercentageChange float64   `json:"total_market_cap_yesterday_percentage_change"`
            TotalVolume24HYesterdayPercentageChange float64   `json:"total_volume_24h_yesterday_percentage_change"`
        } `json:"quote"`
        LastUpdated time.Time `json:"last_updated"`
    } `json:"data"`
    Status Status `json:"status"`
}
