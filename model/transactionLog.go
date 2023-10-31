package model

// Layer - Domain
// 交易紀錄
type TransactionLog struct {
	Id              int    `json:"Id"`
	OperatorID      int    `json:"OperatorID"`
	Content         string `json:"Content"`
	Token           string `json:"Token"`
	IP              string `json:"IP"`
	Location        string `json:"Location"`
	Origin          string `json:"Origin"`
	UserAgent       string `json:"UserAgent"`
	TransactionTime string `json:"TransactionTime"`
}
