package utils

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

//RedisClient Redis客户端
type RedisClient struct {
	prefix string
	*redis.Client
}

/**
Init redis connection pool
初始化redis连接池
*/
func InitRedisPool(host, pwd string, db, timeout, poolSize int, namespace string) *RedisClient {
	// Options 参数配置
	options := &redis.Options{
		Addr:        host,
		Password:    pwd,
		DB:          db,
		IdleTimeout: time.Duration(timeout) * time.Second,
		PoolSize:    poolSize,
	}
	redisClient := RedisClient{
		prefix: namespace,
		Client: redis.NewClient(options),
	}
	// Ping redis before use it to check if there is any error while connection
	// 在使用之前ping一下redis,便于在系统启动的时候检查redis是否正确连接
	if err := redisClient.Ping().Err(); err != nil {
		log.Fatalln("Fail to init redis!")
		return nil
	}
	return &redisClient
}

func CloseRedis(redis *RedisClient) {
	err := redis.Close()
	if err != nil {
		log.Println("Fail to close redis")
	} else {
		log.Println("Success to close redis")
	}
}
