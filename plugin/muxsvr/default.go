package muxsvr

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type IEngine interface {
	Routes(...IRoute) IEngine
	Run(int)
}

type IRoute interface{}

type route struct {
	paths []string
}

type engine struct {
	e          *gin.Engine
	keyOfRoute map[string]IRoute
}

func (e *engine) Routes(routes ...IRoute) IEngine {
	return e
}

func (e engine) Run(port int) {
	e.e.Run(fmt.Sprintf(":%d", port))
}

func NewRoute(method string, paths ...string) IRoute {
	return &route{
		paths: paths,
	}
}

func New(opts ...Option) IEngine {
	e := gin.Default()
	for _, o := range opts {
		o(e)
	}
	return &engine{
		e: e,
	}
}
