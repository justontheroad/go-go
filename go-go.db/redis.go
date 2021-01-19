package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisCtx = context.Background()

//RedisConfig 配置
type RedisConfig struct {
	Host string
	Port int16
	Db   int
}

//RedisClient 客户端
type RedisClient struct {
	RedisConfig
	rdb *redis.Client
}

//GetRdb 获取rdb
func GetRdb(client *RedisClient) {
	if client.rdb == nil || nil != client.rdb.Ping(redisCtx) {
		client.rdb = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", client.RedisConfig.Host, client.RedisConfig.Port),
			Password: "",                    // no password set
			DB:       client.RedisConfig.Db, // use default DB
		})
	}
}

//Get 获取
func (rc *RedisClient) Get(key string) (val string, err error) {
	val, err = rc.rdb.Get(redisCtx, key).Result()
	switch {
	case err == redis.Nil:
		err = nil
		val = ""
	case err != nil:
		log.Println("Get Failed", err)
	case val == "":
		log.Println("value is empty")
	}
	return
}

//Set 设置
func (rc *RedisClient) Set(key string, value string, ttl time.Duration) {
	rc.rdb.Set(redisCtx, key, value, ttl)
}

//IncrAndExpire 自增和设置过期时间，使用pipelie实现
func (rc *RedisClient) IncrAndExpire(key string, ttl time.Duration) (val int64, err error) {
	var incr *redis.IntCmd

	_, err = rc.rdb.Pipelined(redisCtx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(redisCtx, key)
		pipe.Expire(redisCtx, key, ttl)
		return nil
	})

	if err != nil {
		panic(err)
	}

	val = incr.Val()
	fmt.Println(val)
	return
}
