package etherscan

import (
    "encoding/json"
    "net/url"
)

func (c *Configuration) GetTransactions(address string) (*TransactionsResponse, error) {

    values := url.Values{}
    values.Add("module", "account")
    values.Add("action", "txlist")
    values.Add("address", address)
    //values.Add("startblock", startBlock)
    //values.Add("endblock", endBlock)
    //values.Add("page", page)
    //values.Add("offset", offset)
    //values.Add("sort", sort)

    resBody, err := doGetRequest(values, c)
    if err != nil {
        return nil, err
    }

    var etherscanResponse TransactionsResponse

    unmarshalError := json.Unmarshal(resBody.Body, &etherscanResponse)
    if unmarshalError != nil {
        return nil, unmarshalError
    }

    return &etherscanResponse, nil
}

type TransactionsResponse struct {
    Status  string        `json:"status"`
    Message string        `json:"message"`
    Result  []Transaction `json:"result"`
}

type Transaction struct {
    BlockNumber       string `json:"blockNumber"`
    TimeStamp         string `json:"timeStamp"`
    Hash              string `json:"hash"`
    Nonce             string `json:"nonce"`
    BlockHash         string `json:"blockHash"`
    TransactionIndex  string `json:"transactionIndex"`
    From              string `json:"from"`
    To                string `json:"to"`
    Value             string `json:"value"`
    Gas               string `json:"gas"`
    GasPrice          string `json:"gasPrice"`
    IsError           string `json:"isError"`
    TxreceiptStatus   string `json:"txreceipt_status"`
    Input             string `json:"input"`
    ContractAddress   string `json:"contractAddress"`
    CumulativeGasUsed string `json:"cumulativeGasUsed"`
    GasUsed           string `json:"gasUsed"`
    Confirmations     string `json:"confirmations"`
    MethodId          string `json:"methodId"`
    FunctionName      string `json:"functionName"`
}
