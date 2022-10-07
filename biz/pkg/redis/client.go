package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/viper"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	redisv8 "github.com/go-redis/redis/v8"
)

var Client *redisv8.Client
var config *viper.Redis

// InitRedis 初始化Redis
func InitRedis() {
	// 配置初始化
	config = viper.Conf.Redis

	initRedis(context.Background(), &Client)
}

// initRedis 初始化Redis impl
func initRedis(ctx context.Context, client **redisv8.Client) {
	rdb := redisv8.NewClient(&redisv8.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
	})
	if rdb == nil {
		hlog.CtxFatalf(ctx, "[Redis] Init Failed")
	}
	*client = rdb
	hlog.CtxInfof(ctx, "[Redis] PING: %s\n", Client.Ping(ctx))
}

// GetEncodeLock 获取加密操作Redis锁
func GetEncodeLock(ctx context.Context, key string) bool {
	lockSuccess, err := Client.SetNX(ctx, fmt.Sprintf("encode_lock_%v", key), "w", time.Duration(config.EncodeLockSecond)*time.Second).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] GetEncodeLock failed, key: %v, err: %v", key, err)
		return false
	}
	return lockSuccess
}

// GetDecodeLock 获取解密操作Redis锁
func GetDecodeLock(ctx context.Context, key string) bool {
	lockSuccess, err := Client.SetNX(ctx, fmt.Sprintf("decode_lock_%v", key), "w", time.Duration(config.DecodeLockSecond)*time.Second).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] GetDecodeLock failed, key: %v, err: %v", key, err)
		return false
	}
	return lockSuccess
}

// GetIncrId 获取Redis计数器
func GetIncrId(ctx context.Context, key string) int64 {
	id, err := Client.Incr(ctx, key).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] GetIncrId failed, key: %v, err: %v", key, err)
		return 0
	}
	return id
}
