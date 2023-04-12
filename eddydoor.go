package main

import (
	_ "github.com/xm-chentl/eddydoor/internal/api"
	"github.com/xm-chentl/eddydoor/internal/config"
	"github.com/xm-chentl/eddydoor/internal/service/svc"
	"github.com/xm-chentl/eddydoor/plugin/muxsvr"
	"github.com/xm-chentl/eddydoor/plugin/redisex"
	"github.com/xm-chentl/eddydoor/plugin/redisex/goredis"
	"github.com/xm-chentl/eddydoor/utils/guidex"
	"github.com/xm-chentl/eddydoor/utils/guidex/snowflake"
	"github.com/xm-chentl/gocore/iocex"
	"github.com/xm-chentl/goresource"
	"github.com/xm-chentl/goresource/mongoex"
	"github.com/xm-chentl/goresource/mysqlex"
)

func main() {
	config.Init()
	// inject
	iocex.SetMap(new(goresource.IResource), map[string]interface{}{
		"mysql": mysqlex.New(config.Cfg.Resource.MySql.Dsn),
		"mongo": mongoex.New("eddydoor", config.Cfg.Resource.Mongo.Dsn),
	})
	iocex.Set(new(guidex.IGenerate), snowflake.New())
	iocex.Set(new(redisex.IRedis), goredis.New(redisex.Option{
		Addr: config.Cfg.Resource.Redis.Addr,
	}))

	// run
	muxsvr.New(
		svc.NewPost(
			svc.HandlerGetHandler(),
			svc.HandlerOAuthUser(),
			svc.HandlerOAuthAdmin(),
		),
	).Run(config.Cfg.Runtime.Port)
}
