package utils

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func JWTExpiryUnix(tokenString string) (int64, error) {
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())
	claims := jwt.MapClaims{}
	_, _, err := parser.ParseUnverified(tokenString, claims)
	if err != nil {
		return 0, err
	}
	if expAny, ok := claims["exp"]; ok {
		switch v := expAny.(type) {
		case float64:
			return int64(v), nil
		case int64:
			return v, nil
		}
	}
	return 0, errors.New("exp not found")
}

func JWTIsExpired(tokenString string, skewSec int64) bool {
	exp, err := JWTExpiryUnix(tokenString)
	if err != nil { return true }
	now := time.Now().Unix()
	return now+skewSec >= exp
}
