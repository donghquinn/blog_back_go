package types

import "github.com/golang-jwt/jwt/v5"

type JwtInterface struct {
	UserId string `json:"userId"`
	UserEmail string `json:"userEmail"`
	UserType string `json:"userType"`
	Uuid string `json:"uuid"`
	jwt.MapClaims
}

type RedisToken struct {
	TokenItem map[string]string
}