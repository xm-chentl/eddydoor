package ginex

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

type DefaultResp struct {
	Code    int         `json:""`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
