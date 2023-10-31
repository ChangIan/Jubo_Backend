// 中介層 暫時無用
package middleware

import (
	"fmt"

	"github.com/golang-jwt/jwt"

	common "changian.com/jubo/model/common"

	global "changian.com/jubo/startup"
)

// 產生身份辨識碼
func GenerateToken(operatorID int, operatorType int, account string, name string, email string) (string, error) {

	hmacSampleSecret := []byte(global.Config.JwtKey)

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = common.OperatorClaims{
		OperatorID:   operatorID,
		OperatorType: operatorType,
		Account:      account,
		Name:         name,
		Email:        email,
	}

	return token.SignedString(hmacSampleSecret)
}

// 驗證身份辨識碼
func ValidateToken(tokenString string) *common.OperatorClaims {

	hmacSampleSecret := []byte(global.Config.JwtKey)

	token, err := jwt.ParseWithClaims(tokenString, &common.OperatorClaims{}, func(t *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		fmt.Println(err)

		return &common.OperatorClaims{}
	}

	claims := token.Claims.(*common.OperatorClaims)

	return claims
}
