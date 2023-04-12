package usersvc

import "github.com/golang-jwt/jwt/v5"

type LoginUser struct {
	jwt.RegisteredClaims

	ID        string `json:"id"`
	Phone     string `json:"phone"`
	Nickname  string `json:"nickname"`
	LoginTime int64  `json:"login_time"`
}
