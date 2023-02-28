package ginex

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"

	"github.com/xm-chentl/eddydoor/internal/common"
	"github.com/xm-chentl/eddydoor/internal/contract"
	"github.com/xm-chentl/eddydoor/internal/service/usersvc"
	"github.com/xm-chentl/eddydoor/utils/handlers"
	"github.com/xm-chentl/gocore/iocex"
)

type IEngine interface {
	Run(int)
}

type engine struct {
	e *gin.Engine
}

func (e engine) Run(port int) {
	e.e.Run(fmt.Sprintf(":%d", port))
}

func New(opts ...Option) IEngine {
	e := gin.Default()
	for _, o := range opts {
		o(e)
	}
	return engine{
		e: e,
	}
}

func NewDefaultPost() Option {
	return func(e *gin.Engine) {
		validate := validator.New()
		e.POST("/:module/:action", func(ctx *gin.Context) {
			var err error
			resp := DefaultResp{
				Data: RespStruct,
			}
			defer func() {
				if err != nil {
					resp.Code = 500
					resp.Message = err.Error()
					if c, ok := err.(*common.CustomError); ok {
						resp.Code = c.Code
						resp.Message = c.Message
					}
				}
				ctx.JSON(http.StatusOK, resp)
			}()

			route := fmt.Sprintf("/%s/%s", ctx.Param("module"), ctx.Param("action"))
			handler, ok := handlers.Get(route)
			if !ok {
				err = errors.New("找不到接口")
				return
			}

			api := reflect.New(reflect.TypeOf(handler).Elem()).Interface()
			err = ctx.Bind(api)
			if err != nil {
				err = errors.New("参数错误")
				return
			}
			if err = iocex.Inject(api); err != nil {
				err = errors.New("注入失败")
				return
			}
			if _, ok := api.(contract.ISession); ok {
				token := ctx.Request.Header.Get(common.HEADER_Token)
				if token == "" {
					err = errors.New("未登录授权")
					return
				}

				// 获取用户信息
				user := usersvc.LoginUser{}
				tkn, err := jwt.ParseWithClaims(token, &user, func(t *jwt.Token) (interface{}, error) {
					return []byte(viper.GetString("private.secret")), nil
				})
				if err != nil {

					return
				}
				if !tkn.Valid {
					err = errors.New("Invalid token")
					return
				}
				// 赋值ISession
				sessionImpl := &Session{
					User: user,
				}
				srt := reflect.TypeOf(sessionImpl)
				rt := reflect.TypeOf(api).Elem()
				rv := reflect.ValueOf(api).Elem()
				for i := 0; i < rt.NumField(); i++ {
					fs := rt.Field(i)
					if fs.Type.Kind() == reflect.Interface && srt.Implements(fs.Type) {
						rv.Field(i).Set(reflect.ValueOf(sessionImpl))
					}
				}
			}

			tmpHandler, ok := api.(handlers.IHandler)
			if !ok {
				err = errors.New("不是有效的接口")
				return
			}

			err = validate.Struct(tmpHandler)
			if err != nil {
				return
			}

			resp.Data, err = tmpHandler.Call(ctx.Request.Context())

			return
		})
	}
}
