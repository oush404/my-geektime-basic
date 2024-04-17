package ioc

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitRedis() redis.Cmdable {
	fmt.Printf("============>%s\n", viper.GetString("redis.addr"))
	return redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.addr"),
	})
}
