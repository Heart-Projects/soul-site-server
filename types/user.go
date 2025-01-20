package types

import "github.com/golang-jwt/jwt/v4"

type UserJwtClaims struct {
	UserId   uint64   `json:"userId"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	Menus    []string `json:"menus"`
	jwt.RegisteredClaims
}
