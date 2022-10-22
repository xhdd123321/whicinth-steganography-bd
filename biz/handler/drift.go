package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/redis"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// ReceiveDrift 接收一封未知的漂流信
func ReceiveDrift(ctx context.Context, c *app.RequestContext) {
	// API 限流
	remoteIp := utils.RemoteIp(c)
	hlog.CtxInfof(ctx, "Request: %v, remoteIp: %v", string(c.URI().Path()), remoteIp)
	if !redis.GetDriftLock(ctx, remoteIp) {
		hlog.CtxErrorf(ctx, "GetDriftLock failed, remoteIp: %v", remoteIp)
		utils.ResponseError(c, fmt.Sprintf("GetDriftLock failed, remoteIp: %v", remoteIp), nil)
		return
	}
	url, err := redis.ReceiveDrift(ctx)
	if err != nil {
		hlog.CtxErrorf(ctx, "ReceiveDrift failed, err: %v", err)
		utils.ResponseError(c, "ReceiveDrift failed", err)
		return
	}
	// Response
	resp := map[string]interface{}{
		"url": url,
	}
	hlog.CtxInfof(ctx, "ReceiveDrift Success, resp: %+v", resp)
	utils.ResponseOK(c, "ReceiveDrift Success", resp)
}
