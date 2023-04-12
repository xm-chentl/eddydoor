package svc

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xm-chentl/eddydoor/internal/common"
	"github.com/xm-chentl/eddydoor/internal/config"
	"github.com/xm-chentl/eddydoor/internal/contract"
	"github.com/xm-chentl/eddydoor/internal/service/adminsvc"
	"github.com/xm-chentl/eddydoor/internal/service/usersvc"
	"github.com/xm-chentl/eddydoor/plugin/muxsvr"
	"github.com/xm-chentl/eddydoor/plugin/muxsvr/ctxtype"
	"github.com/xm-chentl/gocore/iocex"
)

func HandlerGetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				ctx.Set(ctxtype.Err.String(), err)
			}
			ctx.Next()
		}()

		route := fmt.Sprintf("/%s/%s", ctx.Param("module"), ctx.Param("action"))
		handler, ok := muxsvr.GetHandler(route)
		if !ok {
			err = muxsvr.ErrAPIHandlerNotExist
			return
		}
		// 找到逻辑代码块
		api := reflect.New(reflect.TypeOf(handler).Elem()).Interface()
		err = ctx.Bind(api)
		if err != nil {
			err = muxsvr.ErrHandlerParameterFailed
			return
		}
		if err = iocex.Inject(api); err != nil {
			log.Println(err)
			err = muxsvr.ErrHandlerInjectFailed
			return
		}
		ctx.Set(ctxtype.Handler.String(), api)
	}
}

func HandlerOAuthUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		if errImp, ok := ctx.Get(ctxtype.Err.String()); ok && errImp != nil {
			err = errImp.(error)
			return
		}
		defer func() {
			if err != nil {
				ctx.Set(ctxtype.Err.String(), err)
			}
			ctx.Next()
		}()

		api, _ := ctx.Get(ctxtype.Handler.String())
		if _, ok := api.(contract.IUserSession); ok {
			token := ctx.Request.Header.Get(common.HEADER_Token)
			if token == "" {
				err = errors.New("未登录授权")
				return
			}

			// 获取用户信息
			user := usersvc.LoginUser{}
			tkn, err := jwt.ParseWithClaims(token, &user, func(t *jwt.Token) (interface{}, error) {
				return []byte(config.Cfg.Secret.User), nil
			})
			if err != nil {

				return
			}
			if !tkn.Valid || user.ID == "" {
				err = errors.New("Invalid token")
				return
			}
			// 赋值ISession
			sessionImpl := &UserSession{
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
			ctx.Set(ctxtype.Handler.String(), api)
		}
	}
}

func HandlerOAuthAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		if errImp, ok := ctx.Get(ctxtype.Err.String()); ok && errImp != nil {
			err = errImp.(error)
			return
		}
		defer func() {
			if err != nil {
				ctx.Set(ctxtype.Err.String(), err)
			}
			ctx.Next()
		}()

		api, _ := ctx.Get(ctxtype.Handler.String())
		if _, ok := api.(contract.IAdminSession); ok {
			token := ctx.Request.Header.Get(common.HEADER_ADMIN_TOKEN)
			if token == "" {
				err = errors.New("此管理员未登录授权")
				return
			}

			// 获取用户信息
			admin := adminsvc.LoginAdmin{}
			tkn, err := jwt.ParseWithClaims(token, &admin, func(t *jwt.Token) (interface{}, error) {
				return []byte(config.Cfg.Secret.Admin), nil
			})
			fmt.Println("data: ", admin)
			if err != nil {
				log.Println("jwt err: ", err.Error())
				return
			}
			if !tkn.Valid || admin.ID == "" {
				err = errors.New("Invalid token")
				return
			}
			// 赋值ISession
			sessionImpl := &AdminSession{
				Admin: admin,
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
			ctx.Set(string(ctxtype.Handler), api)
		}
	}
}

func NewPost(handlers ...gin.HandlerFunc) muxsvr.Option {
	return func(e *gin.Engine) {
		validate := validator.New()
		for _, handler := range handlers {
			e.Use(handler)
		}
		e.POST("/:module/:action", func(ctx *gin.Context) {
			var err error
			resp := muxsvr.DefaultResp{
				Code: 200,
				Data: muxsvr.RespStruct,
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

			// // 识别路由
			// route := fmt.Sprintf("/%s/%s", ctx.Param("module"), ctx.Param("action"))
			// handler, ok := muxsvr.GetHandler(route)
			// if !ok {
			// 	err = muxsvr.ErrAPIHandlerNotExist
			// 	return
			// }

			// // 找到逻辑代码块
			// api := reflect.New(reflect.TypeOf(handler).Elem()).Interface()
			// err = ctx.Bind(api)
			// if err != nil {
			// 	err = muxsvr.ErrHandlerParameterFailed
			// 	return
			// }
			// if err = iocex.Inject(api); err != nil {
			// 	err = errors.New("注入失败")
			// 	return
			// }
			// // 鉴权判断
			// if _, ok := api.(contract.IUserSession); ok {
			// 	token := ctx.Request.Header.Get(common.HEADER_Token)
			// 	if token == "" {
			// 		err = errors.New("未登录授权")
			// 		return
			// 	}

			// 	// 获取用户信息
			// 	user := usersvc.LoginUser{}
			// 	tkn, err := jwt.ParseWithClaims(token, &user, func(t *jwt.Token) (interface{}, error) {
			// 		return []byte(config.Cfg.Secret.User), nil
			// 	})
			// 	if err != nil {

			// 		return
			// 	}
			// 	if !tkn.Valid {
			// 		err = errors.New("Invalid token")
			// 		return
			// 	}
			// 	// 赋值ISession
			// 	sessionImpl := &UserSession{
			// 		User: user,
			// 	}
			// 	srt := reflect.TypeOf(sessionImpl)
			// 	rt := reflect.TypeOf(api).Elem()
			// 	rv := reflect.ValueOf(api).Elem()
			// 	for i := 0; i < rt.NumField(); i++ {
			// 		fs := rt.Field(i)
			// 		if fs.Type.Kind() == reflect.Interface && srt.Implements(fs.Type) {
			// 			rv.Field(i).Set(reflect.ValueOf(sessionImpl))
			// 		}
			// 	}
			// }
			// if _, ok := api.(contract.IAdminSession); ok {
			// 	token := ctx.Request.Header.Get(common.HEADER_ADMIN_TOKEN)
			// 	if token == "" {
			// 		err = errors.New("此管理员未登录授权")
			// 		return
			// 	}

			// 	// 获取用户信息
			// 	admin := adminsvc.LoginAdmin{}
			// 	tkn, err := jwt.ParseWithClaims(token, &admin, func(t *jwt.Token) (interface{}, error) {
			// 		return []byte(config.Cfg.Secret.Admin), nil
			// 	})
			// 	fmt.Println("data: ", admin)
			// 	if err != nil {
			// 		log.Println("jwt err: ", err.Error())
			// 		return
			// 	}
			// 	if !tkn.Valid {
			// 		err = errors.New("Invalid token")
			// 		return
			// 	}
			// 	// 赋值ISession
			// 	sessionImpl := &AdminSession{
			// 		Admin: admin,
			// 	}
			// 	srt := reflect.TypeOf(sessionImpl)
			// 	rt := reflect.TypeOf(api).Elem()
			// 	rv := reflect.ValueOf(api).Elem()
			// 	for i := 0; i < rt.NumField(); i++ {
			// 		fs := rt.Field(i)
			// 		if fs.Type.Kind() == reflect.Interface && srt.Implements(fs.Type) {
			// 			rv.Field(i).Set(reflect.ValueOf(sessionImpl))
			// 		}
			// 	}
			// }
			if errImp, ok := ctx.Get(ctxtype.Err.String()); ok && errImp != nil {
				err = errImp.(error)
				return
			}

			// 处理函数判断
			handler, _ := ctx.Get(ctxtype.Handler.String())
			apiHandler, _ := handler.(muxsvr.IHandler)
			err = validate.Struct(apiHandler)
			if err != nil {
				return
			}
			resp.Data, err = apiHandler.Call(ctx)

			return
		})
	}
}
