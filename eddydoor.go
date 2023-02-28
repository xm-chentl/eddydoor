package main

import (
	"github.com/spf13/viper"
	_ "github.com/xm-chentl/eddydoor/internal/api"
	"github.com/xm-chentl/eddydoor/utils/ginex"
	"github.com/xm-chentl/eddydoor/utils/guidex"
	"github.com/xm-chentl/eddydoor/utils/guidex/snowflake"
	"github.com/xm-chentl/eddydoor/utils/redisex"
	"github.com/xm-chentl/eddydoor/utils/redisex/goredis"
	"github.com/xm-chentl/gocore/iocex"
	"github.com/xm-chentl/goresource"
	"github.com/xm-chentl/goresource/mysqlex"
)

func main() {
	// config
	viper.SetConfigFile("config/config.toml")
	if err := viper.ReadInConfig(); err != nil {
		panic("reading config failed: " + err.Error())
	}

	// inject
	iocex.Set(new(goresource.IResource), mysqlex.New(viper.GetString("resource.mysql.dsn")))
	iocex.Set(new(guidex.IGenerate), snowflake.New())
	iocex.Set(new(redisex.IRedis), goredis.New(redisex.Option{
		Addr: viper.GetString("resource.redis.addr"),
	}))

	// run
	ginex.New(
		ginex.NewDefaultPost(),
	).Run(viper.GetInt("runtime.port"))
}
