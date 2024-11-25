package main

import (
	"context"

	"github.com/gogf/gf/os/glog"
	"github.com/redis/go-redis/v9"
)

const (
	RedisUrl = "redis://localhost:6379/0"
)

var logger *glog.Logger

var rdb *redis.Client

func init() {
	logger = glog.New()
	logger.SetFlags(glog.F_TIME_STD | glog.F_FILE_SHORT)
	logger.Info("logger initial success")

	options, err := redis.ParseURL(RedisUrl)
	if err != nil {
		logger.Fatal(err)
	}

	rdb = redis.NewClient(options)

	ctx := context.Background()
	result := rdb.Ping(ctx)
	if result.Err() != nil {
		logger.Errorf("redis ping error: %v", result.Err())
		return
	}
}

func main() {
	logger.Info("service start success")
}
