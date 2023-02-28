package user

import (
	"context"
)

type LoginAPI struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (s LoginAPI) Call(ctx context.Context) (res interface{}, err error) {
	return
}
