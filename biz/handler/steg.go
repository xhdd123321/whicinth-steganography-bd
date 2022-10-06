package handler

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/redis"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/service/fileService"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/service/stegService"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"
)

// EncodeImageFromImage 图片中加密图片
func EncodeImageFromImage(ctx context.Context, c *app.RequestContext) {
	// 获取文件夹唯一标识
	key := "encode_image"
	id := redis.GetIncrId(ctx, key)
	if id <= 0 {
		utils.ResponseError(c, "GetIncrId failed, key: encode_image", nil)
		return
	}
	dir := filepath.Join(".", "media", key, fmt.Sprintf("%v", id))
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		utils.ResponseError(c, "MkdirAll failed", err)
		return
	}
	// 上传载体文件
	carrierUploadPath := filepath.Join(dir, "carrier.png")
	if err := fileService.UploadFile(ctx, c, "carrier_file", carrierUploadPath); err != nil {
		logger.CtxErrorf(ctx, "UploadFile failed, path: %v, err: %v", carrierUploadPath, err)
		utils.ResponseError(c, "UploadFile failed", err)
		return
	}
	// 上传数据文件
	dataUploadPath := filepath.Join(dir, "data.png")
	if err := fileService.UploadFile(ctx, c, "data_file", dataUploadPath); err != nil {
		logger.CtxErrorf(ctx, "UploadFile failed, path: %v, err: %v", dataUploadPath, err)
		utils.ResponseError(c, "UploadFile failed", err)
		return
	}
	// 获取加密文件
	resultFilePath := filepath.Join(dir, "result.png")
	err = stegService.EncodeImage(carrierUploadPath, dataUploadPath, resultFilePath)
	if err != nil {
		logger.CtxErrorf(ctx, "EncodeImage failed, path: %v, err: %v", resultFilePath, err)
		utils.ResponseError(c, "EncodeImage failed", err)
		return
	}
	resp := map[string]interface{}{
		"carrier_file": carrierUploadPath,
		"data_file":    dataUploadPath,
		"result_file":  resultFilePath,
	}
	utils.ResponseOK(c, "EncodeImageFromImage Success", resp)
}

// DecodeImageFromImage 图片中解密图片
func DecodeImageFromImage(ctx context.Context, c *app.RequestContext) {
}

// EncodeDocFromImage 图片中加密文字
func EncodeDocFromImage(ctx context.Context, c *app.RequestContext) {
}

// DecodeDocFromImage 图片中解密文字
func DecodeDocFromImage(ctx context.Context, c *app.RequestContext) {
}
