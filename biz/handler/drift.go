package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/redis"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// ReceiveDrift 接收一封未知的漂流信
func ReceiveDrift(ctx context.Context, c *app.RequestContext) {
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
	utils.ResponseOK(c, "ReceiveDrift Success", resp)
}
