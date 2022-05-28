package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var nowDate = time.Now().Format("2006-01-02 15")
var secret = fmt.Sprintf("%v%v", nowDate, "xxxx")

// GenerateToken 生成Token值
func GenerateToken(mapClaims jwt.MapClaims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(key))
}

// token: "eyJhbGciO...解析token"
func ParseToken(token string, secret string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["cmd"].(string), nil
}

func main() {
	dict := make(map[string]interface{})
	dict["name"] = "xxxx"
	dict["age"] = "x18"

	dict1 := make(map[string]interface{})
	dict1["name"] = "xxxxx"
	dict1["age"] = "18"
	tokenNew, _ := GenerateToken(dict, "")   // 生成token
	tokenNew1, _ := GenerateToken(dict1, "") // 生成token
	fmt.Printf("%s \n %s", tokenNew, tokenNew1)
}
