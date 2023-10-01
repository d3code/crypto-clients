package cmc

import (
    "encoding/json"
    "net/url"
    "time"
)

func (c *Configuration) GetBlockchainStatisticsLatest(q url.Values) (*BlockchainStatisticsLatest, error) {
    requestURL := "/v1/blockchain/statistics/latest"

    response, err := request(requestURL, q, c)
    if err != nil {
        return nil, err
    }

    var blockchainStatisticsLatest BlockchainStatisticsLatest
    unmarshalError := json.Unmarshal(response.Body, &blockchainStatisticsLatest)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &blockchainStatisticsLatest, nil
}

type BlockchainStatisticsLatest struct {
    Data map[string]struct {
        Id                  int       `json:"id"`
        Slug                string    `json:"slug"`
        Symbol              string    `json:"symbol"`
        BlockRewardStatic   float64   `json:"block_reward_static"`
        ConsensusMechanism  string    `json:"consensus_mechanism"`
        Difficulty          string    `json:"difficulty"`
        Hashrate24H         string    `json:"hashrate_24h"`
        PendingTransactions int       `json:"pending_transactions"`
        ReductionRate       string    `json:"reduction_rate"`
        TotalBlocks         int       `json:"total_blocks"`
        TotalTransactions   string    `json:"total_transactions"`
        Tps24H              float64   `json:"tps_24h"`
        FirstBlockTimestamp time.Time `json:"first_block_timestamp"`
    } `json:"data"`
    Status Status `json:"status"`
}
