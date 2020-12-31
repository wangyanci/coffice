package model

import (
	"github.com/dgrijalva/jwt-go"
)

type ServerClaims struct {
	DomainID   string `json:"userId"`
	DomainName string `json:"userName"`
	Secret     string `json:"password"`
	jwt.StandardClaims
}
