package app

import (
	"context"

	"github.com/xm-chentl/eddydoor/internal/contract"
)

type GetAPI struct {
	contract.IUserSession
}

func (s GetAPI) Call(ctx context.Context) (res interface{}, err error) {
	res = s.GetUser()
	return
}
