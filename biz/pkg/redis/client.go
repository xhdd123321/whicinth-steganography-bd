package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/viper"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	redisv8 "github.com/go-redis/redis/v8"
)

var (
	Client *redisv8.Client
	config *viper.Redis
)

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

// GetValue 获取redis value
func GetValue(ctx context.Context, key string) string {
	res, err := Client.Get(ctx, key).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] GetIncrId failed, key: %v, err: %v", key, err)
		return ""
	}
	return res
}

const DRIFT_KEY = "drift"

// AddDriftSet 漂流信集合新增元素
func AddDriftSet(ctx context.Context, value string) bool {
	key := DRIFT_KEY
	res, err := Client.SAdd(ctx, key, value).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] AppendDriftSet failed, key: %v, value: %v, err: %v", key, value, err)
		return false
	}
	// 移除超过限制的历史元素
	if res > int64(config.DriftLimit) {
		_, err = Client.SPopN(ctx, key, int64(config.DriftLimit)-res).Result()
		if err != nil {
			hlog.CtxErrorf(ctx, "[Redis] DriftSet SPopN failed, key: %v, value: %v, err: %v", key, err)
		}
	}
	return true
}

// ReceiveDrift 接收漂流信
func ReceiveDrift(ctx context.Context) (string, error) {
	key := DRIFT_KEY
	drift, err := Client.SRandMember(ctx, key).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] ReceiveDrift failed, key: %v, err: %v", key, err)
		return "", err
	}
	return drift, nil
}

// GetDriftCount 获取漂流信数量
func GetDriftCount(ctx context.Context) int64 {
	key := DRIFT_KEY
	count, err := Client.SCard(ctx, key).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] GetDriftCount failed, key: %v, err: %v", key, err)
		return 0
	}
	return count
}
