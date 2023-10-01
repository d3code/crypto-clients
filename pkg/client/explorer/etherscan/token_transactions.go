package etherscan

import (
    "encoding/json"
    "net/url"
)

func (c *Configuration) GetTokenTransactions(address string, contractAddress string) (*TokenTransactionsResponse, error) {

    values := url.Values{}
    values.Add("module", "account")
    values.Add("action", "tokentx")
    values.Add("contractaddress", contractAddress)
    values.Add("address", address)
    //values.Add("page", page)
    //values.Add("offset", offset)
    //values.Add("startblock", startBlock)
    //values.Add("endblock", endBlock)
    //values.Add("sort", sort)

    resBody, err := doGetRequest(values, c)
    if err != nil {
        return nil, err
    }

    var etherscanResponse TokenTransactionsResponse

    unmarshalError := json.Unmarshal(resBody.Body, &etherscanResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &etherscanResponse, nil
}

func (c *Configuration) GetTokenTransactionsAll(address string) (*TokenTransactionsResponse, error) {

    values := url.Values{}
    values.Add("module", "account")
    values.Add("action", "tokentx")
    values.Add("address", address)
    //values.Add("page", page)
    //values.Add("offset", offset)
    //values.Add("startblock", startBlock)
    //values.Add("endblock", endBlock)
    //values.Add("sort", "asc")

    resBody, err := doGetRequest(values, c)
    if err != nil {
        return nil, err
    }

    var etherscanResponse TokenTransactionsResponse

    unmarshalError := json.Unmarshal(resBody.Body, &etherscanResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &etherscanResponse, nil
}

type TokenTransactionsResponse struct {
    Status  string             `json:"status"`
    Message string             `json:"message"`
    Result  []TokenTransaction `json:"result"`
}

type TokenTransaction struct {
    BlockNumber       string `json:"blockNumber"`
    TimeStamp         string `json:"timeStamp"`
    Hash              string `json:"hash"`
    Nonce             string `json:"nonce"`
    BlockHash         string `json:"blockHash"`
    From              string `json:"from"`
    ContractAddress   string `json:"contractAddress"`
    To                string `json:"to"`
    Value             string `json:"value"`
    TokenName         string `json:"tokenName"`
    TokenSymbol       string `json:"tokenSymbol"`
    TokenDecimal      string `json:"tokenDecimal"`
    TransactionIndex  string `json:"transactionIndex"`
    Gas               string `json:"gas"`
    GasPrice          string `json:"gasPrice"`
    GasUsed           string `json:"gasUsed"`
    CumulativeGasUsed string `json:"cumulativeGasUsed"`
    Input             string `json:"input"`
    Confirmations     string `json:"confirmations"`
}
