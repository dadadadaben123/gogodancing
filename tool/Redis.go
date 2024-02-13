package tool

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"log"
	"time"
)

type RedisStore struct {
	client *redis.Client
}

var Redis RedisStore
var ctx = context.Background()

func InitRedisStore(config *Config) *RedisStore {
	redisCfg := config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.Db,
	})

	Redis.client = client
	base64Captcha.SetCustomStore(&Redis)

	return &Redis

}

// get
func (rs *RedisStore) Get(id string, clear bool) (value string) {
	val, err := rs.client.Get(ctx, id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		//err := rs.client.Del(ctx, id).Err()
		//fmt.Println("clear------------")
		//if err != nil {
		//	log.Println(err)
		//	return ""
		//}
	}
	return val
}

// set
func (rs *RedisStore) Set(id string, value string) {
	fmt.Println("set captcha", id, value)
	err := rs.client.Set(ctx, id, value, time.Minute*100).Err()
	if err != nil {
		log.Println(err)
	}
}
