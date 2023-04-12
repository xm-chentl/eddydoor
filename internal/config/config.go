package config

import (
	"time"

	"github.com/spf13/viper"
)

var (
	Cfg *Config
)

type Config struct {
	Secret   Secret
	Resource Resource
	Runtime  Runtime
}

type Secret struct {
	Admin   string
	User    string
	Expired time.Duration
}

type Resource struct {
	MySql ResourceMySql
	Mongo ResourceMongo
	Redis ResourceRedis
}

type ResourceMySql struct {
	Dsn string
}
type ResourceMongo struct {
	Dsn string
}

type ResourceRedis struct {
	Addr string
}

type Runtime struct {
	Server string
	Port   int
}

func Init() {
	// config
	viper.SetConfigFile("config.toml")
	if err := viper.ReadInConfig(); err != nil {
		panic("reading config failed: " + err.Error())
	}

	Cfg = &Config{
		Secret: Secret{
			Admin:   viper.GetString("secret.admin"),
			User:    viper.GetString("secret.user"),
			Expired: time.Duration(viper.GetInt("secret.loginExpired")),
		},
		Resource: Resource{
			MySql: ResourceMySql{
				Dsn: viper.GetString("resource.mysql.dsn"),
			},
			Mongo: ResourceMongo{
				Dsn: viper.GetString("resource.mongodb.dsn"),
			},
			Redis: ResourceRedis{
				Addr: viper.GetString("resource.redis.addr"),
			},
		},
		Runtime: Runtime{
			Server: viper.GetString("runtime.server"),
			Port:   viper.GetInt("runtime.port"),
		},
	}
}
