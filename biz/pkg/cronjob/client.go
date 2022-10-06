package cronjob

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/service/fileService"
)

// TTL_MINUTE 文件夹最长生存时间
const TTL_MINUTE = 1

func InitCronjob() {
	go expiredFileClearCronJob(context.Background())
}

func expiredFileClearCronJob(ctx context.Context) {
	sched := time.Tick(time.Minute * TTL_MINUTE)
	for range sched {
		err := fileService.ClearFile(ctx, TTL_MINUTE)
		if err != nil {
			hlog.CtxErrorf(ctx, "ClearFile cronjob run failed, err: %v", err)
		}
	}
}
