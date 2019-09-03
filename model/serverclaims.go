package model

import (
	"github.com/dgrijalva/jwt-go"
)

type ServerClaims struct {
	DomainID   string `json:"domain_id"`
	DomainName string `json:"domain_name"`
	Secret     string `json:"Secret"`
	jwt.StandardClaims
}
