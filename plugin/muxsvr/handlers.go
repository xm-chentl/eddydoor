package muxsvr

import "context"

type IHandler interface {
	Call(context.Context) (interface{}, error)
}

var handlerPool = make(map[string]IHandler)

type Handler map[string]IHandler

func GetHandler(key string) (handler IHandler, ok bool) {
	handler, ok = handlerPool[key]
	if ok {
		return
	}

	return
}

func RegisterHandlers(handlers ...Handler) {
	for _, item := range handlers {
		for k, v := range item {
			handlerPool[k] = v
		}
	}
}
