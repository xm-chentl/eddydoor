package admin

import (
	"context"

	"github.com/xm-chentl/eddydoor/internal/contract"
)

type InfoAPI struct {
	contract.IAdminSession
}

func (s InfoAPI) Call(ctx context.Context) (res interface{}, err error) {
	res = map[string]interface{}{
		"user": s.GetAdmin().ToResponse(),
	}
	return
}
