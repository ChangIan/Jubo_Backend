package modelCommon

import "github.com/golang-jwt/jwt"

// Layer - Domain
// 操作者身份聲明
type OperatorClaims struct {
	*jwt.StandardClaims
	OperatorID   int    `json:"OperatorID"`
	OperatorType int    `json:"OperatorType"`
	Account      string `json:"Account"`
	Name         string `json:"Name"`
	Phone        string `json:"Phone"`
	Email        string `json:"Email"`
	Gender       string `json:"Gender"`
}
