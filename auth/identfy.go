package auth

import (
	"fmt"
	"time"

	"vueApp/model"
	"vueApp/setting"

	"github.com/dgrijalva/jwt-go"
)

func GetAuthToken(user model.User) (string, error){
	claims := model.ServerClaims{
		Secret: user.Secret,
		DomainID: user.DomainId,
		DomainName: user.DomainName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(setting.SigningKey))
}

func Keyfunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("not authorization")
	}
	return []byte(setting.SigningKey), nil
}


