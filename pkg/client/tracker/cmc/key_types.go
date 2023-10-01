package cmc

import "time"

type KeyInfo struct {
    Data struct {
        Plan struct {
            CreditLimitDaily                 int       `json:"credit_limit_daily"`
            CreditLimitDailyReset            string    `json:"credit_limit_daily_reset"`
            CreditLimitDailyResetTimestamp   time.Time `json:"credit_limit_daily_reset_timestamp"`
            CreditLimitMonthly               int       `json:"credit_limit_monthly"`
            CreditLimitMonthlyReset          string    `json:"credit_limit_monthly_reset"`
            CreditLimitMonthlyResetTimestamp time.Time `json:"credit_limit_monthly_reset_timestamp"`
            RateLimitMinute                  int       `json:"rate_limit_minute"`
        } `json:"plan"`
        Usage struct {
            CurrentMinute keyRequests `json:"current_minute"`
            CurrentDay    keyRequests `json:"current_day"`
            CurrentMonth  keyRequests `json:"current_month"`
        } `json:"usage"`
    } `json:"data"`
    Status Status `json:"status"`
}

type keyRequests struct {
    RequestsMade int `json:"requests_made"`
    RequestsLeft int `json:"requests_left"`
}
