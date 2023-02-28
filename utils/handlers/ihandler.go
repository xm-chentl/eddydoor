package handlers

import "context"

type IHandler interface {
	Call(context.Context) (interface{}, error)
}
