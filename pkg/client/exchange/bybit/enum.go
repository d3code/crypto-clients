package bybit

type AccountType string

const (
    CONTRACT   AccountType = "CONTRACT"   // Derivatives Account
    SPOT       AccountType = "SPOT"       // Spot Account
    INVESTMENT AccountType = "INVESTMENT" // ByFi Account
    OPTION     AccountType = "OPTION"     // USDC Account
    UNIFIED    AccountType = "UNIFIED"    // UMA or UTA
    FUND       AccountType = "FUND"       // Funding Account
)
