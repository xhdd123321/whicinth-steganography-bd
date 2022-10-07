package cronjob

import (
	"context"
	"math/rand"
	"time"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/viper"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/qiniu"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/service/fileService"
)

var config *viper.Cronjob

// InitCronjob 初始化定时任务
func InitCronjob() {
	// 配置初始化
	config = viper.Conf.Cronjob

	ctx := context.Background()
	go expiredFileClearCronJob(ctx)
	hlog.CtxInfof(ctx, "[Cronjob] ExpiredFileClearCronJob start...")
	go refreshUploadTokenCronJob(ctx)
	hlog.CtxInfof(ctx, "[Cronjob] RefreshUploadTokenCronJob start...")
}

// expiredFileClearCronJob 过期文件清理任务
func expiredFileClearCronJob(ctx context.Context) {
	// 启动时先执行一遍清理
	err := fileService.ClearFile(ctx, config.TempFileMinute)
	if err != nil {
		hlog.CtxErrorf(ctx, "ClearFile cronjob run failed, err: %v", err)
	}
	// 启动定时清理任务
	sched := time.Tick(time.Minute * time.Duration(config.TempFileMinute))
	for range sched {
		err := fileService.ClearFile(ctx, config.TempFileMinute)
		if err != nil {
			hlog.CtxErrorf(ctx, "ClearFile cronjob run failed, err: %v", err)
		}
	}
}

// refreshUploadTokenCronJob 对象存储上传凭证定时任务
func refreshUploadTokenCronJob(ctx context.Context) {
	// 启动定时刷新uploadToken任务
	sched := time.Tick(time.Minute * time.Duration(config.TokenMinute))
	for range sched {
		qiniu.RefreshToken()
		newToken := qiniu.GetUpToken()
		// Token获取失败进行一次退避重试(5~15s)
		if newToken == "" {
			hlog.CtxErrorf(ctx, "RefreshToken failed, let's retry")
			time.Sleep(time.Second * time.Duration(5+rand.Intn(10)))
			qiniu.RefreshToken()
			hlog.CtxErrorf(ctx, "Retry result token %v", qiniu.GetUpToken())
		} else {
			hlog.CtxInfof(ctx, "RefreshToken success, newToken: %v", newToken)
		}
	}
}
