package user

import (
	"context"

	"github.com/xm-chentl/eddydoor/internal/contract"
)

type RefreshTokenAPI struct {
	contract.ISession
}

func (s RefreshTokenAPI) Call(ctx context.Context) (res interface{}, err error) {
	return
}
