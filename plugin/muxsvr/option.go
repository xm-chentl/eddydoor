package muxsvr

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

type DefaultResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
