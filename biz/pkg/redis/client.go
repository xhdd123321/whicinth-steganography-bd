package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	redisv8 "github.com/go-redis/redis/v8"
)

var Client *redisv8.Client

func InitRedis() {
	initRedis(context.Background(), &Client)
}

func initRedis(ctx context.Context, client **redisv8.Client) {
	rdb := redisv8.NewClient(&redisv8.Options{
		Addr:     "123.57.150.240:6379",
		Password: "Fblr9ovO",
		DB:       9,
	})
	if rdb == nil {
		hlog.CtxFatalf(ctx, "[Redis] Init Failed")
	}
	*client = rdb
	hlog.CtxInfof(ctx, "[Redis] PING: %s\n", Client.Ping(ctx))
}

func GetEncodeLock(ctx context.Context, key string) bool {
	lockSuccess, err := Client.SetNX(ctx, fmt.Sprintf("encode_lock_%v", key), "w", 10*time.Second).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] GetEncodeLock failed, err: %v", err)
		return false
	}
	return lockSuccess
}

func GetIncrId(ctx context.Context, key string) int64 {
	id, err := Client.Incr(ctx, key).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] GetEncodeLock failed, err: %v", err)
		return 0
	}
	return id
}
