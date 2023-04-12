package user

import "context"

// CheckAPI is 检查是否有效
type CheckAPI struct{}

func (s CheckAPI) Call(ctx context.Context) (res interface{}, err error) {
	return
}
