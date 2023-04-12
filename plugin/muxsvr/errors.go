package muxsvr

import "errors"

var (
	ErrAPIHandlerNotExist     = errors.New("api handler is not exists")
	ErrHandlerParameterFailed = errors.New("api handler parameter failed")
	ErrHandlerInjectFailed    = errors.New("api handler inject failed")
)
