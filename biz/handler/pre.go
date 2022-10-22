package handler

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/service/fileService"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/redis"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"
)

const (
	PRE_FILE           = "pre_file"
	COMPRESS_IMAGE_KEY = "compress_image"
)

// CompressImageByTinify 使用Tinify压缩图片
func CompressImageByTinify(ctx context.Context, c *app.RequestContext) {
	// API 限流
	remoteIp := utils.RemoteIp(c)
	hlog.CtxInfof(ctx, "Request: %v, remoteIp: %v", string(c.URI().Path()), remoteIp)
	if !redis.GetCompressLock(ctx, remoteIp) {
		hlog.CtxErrorf(ctx, "GetCompressLock failed, remoteIp: %v", remoteIp)
		utils.ResponseError(c, fmt.Sprintf("GetCompressLock failed, remoteIp: %v", remoteIp), nil)
		return
	}
	// 创建临时文件夹
	key := COMPRESS_IMAGE_KEY
	id := redis.GetIncrId(ctx, key)
	if id <= 0 {
		hlog.CtxErrorf(ctx, "GetIncrId failed, key: %v", key)
		utils.ResponseError(c, fmt.Sprintf("GetIncrId failed, key: %v", key), nil)
		return
	}
	dir := filepath.Join(".", "media", key, fmt.Sprintf("%v", id))
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		hlog.CtxErrorf(ctx, "MkdirAll failed, err: %v", err)
		utils.ResponseError(c, "MkdirAll failed", err)
		return
	}
	// 清理临时文件夹
	defer func(path string) {
		err = os.RemoveAll(path)
		hlog.CtxErrorf(ctx, "Clear temp file failed, path: %v, err: %v", path, err)
	}(dir)
	// 上传预处理文件
	preUploadPath := filepath.Join(dir, "pre.png")
	if err := fileService.UploadFile(ctx, c, PRE_FILE, preUploadPath); err != nil {
		hlog.CtxErrorf(ctx, "UploadFile failed, path: %v, err: %v", preUploadPath, err)
		utils.ResponseError(c, "UploadFile failed", err)
		return
	}
	// 压缩图片
	res, err := fileService.CompressImage(ctx, preUploadPath)
	if err != nil {
		hlog.CtxErrorf(ctx, "CompressImageByTinify failed, err: %v", err)
		utils.ResponseError(c, "CompressImageByTinify failed", err)
		return
	}
	// Response
	resp := res
	hlog.CtxInfof(ctx, "CompressImageByTinify Success, resp: %+v", resp)
	utils.ResponseOK(c, "CompressImageByTinify Success", resp)
}
