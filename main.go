package main

import (
	"context"
	"time"

	"github.com/gogf/gf/os/glog"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

const (
	RedisUrl = "redis://localhost:6379/0"
	MysqlUrl = "root:Admin123@tcp(localhost:3306)/demo_db"
)

var logger *glog.Logger

var rdb *redis.Client

var db *sqlx.DB

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

	db, err = sqlx.Open("mysql", MysqlUrl)
	if err != nil {
		logger.Fatal(err)
	}
	db.SetConnMaxLifetime(30 * time.Second)
	logger.Debugf("db initial success")
}

func main() {
	logger.Info("service start success")
}
