package usersvc

import "github.com/dgrijalva/jwt-go"

type LoginUser struct {
	jwt.StandardClaims

	ID        string `json:"id"`
	Phone     string `json:"phone"`
	Nickname  string `json:"nickname"`
	LoginTime int64  `json:"login_time"`
}
