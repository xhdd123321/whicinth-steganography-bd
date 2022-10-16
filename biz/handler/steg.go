package handler

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/qiniu"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/redis"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/service/fileService"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/service/stegService"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"
)

const (
	CARRIER_FILE = "carrier_file"
	DATA_FILE    = "data_file"
	DATA_DOC     = "data_doc"
)

// EncodeImageFromImage 图片中加密图片
func EncodeImageFromImage(ctx context.Context, c *app.RequestContext) {
	// API限流
	remoteIp := utils.RemoteIp(c)
	hlog.CtxInfof(ctx, "Request: %v, remoteIp: %v", string(c.URI().Path()), remoteIp)
	if !redis.GetEncodeLock(ctx, remoteIp) {
		hlog.CtxErrorf(ctx, "GetEncodeLock failed, remoteIp: %v", remoteIp)
		utils.ResponseError(c, fmt.Sprintf("GetEncodeLock failed, remoteIp: %v", remoteIp), nil)
		return
	}
	// 创建临时文件夹
	key := "encode_image"
	id := redis.GetIncrId(ctx, key)
	if id <= 0 {
		hlog.CtxErrorf(ctx, "GetIncrId failed, key: %v", key)
		utils.ResponseError(c, fmt.Sprintf("GetIncrId failed, key: %v", key), nil)
		return
	}
	dir := filepath.Join(".", "media", key, fmt.Sprintf("%v", id))
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		hlog.CtxErrorf(ctx, "MkdirAll failed, err:%v", err)
		utils.ResponseError(c, "MkdirAll failed", err)
		return
	}
	// 清理临时文件夹
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dir)
	// 上传载体文件
	carrierUploadPath := filepath.Join(dir, "carrier.png")
	if err := fileService.UploadFile(ctx, c, CARRIER_FILE, carrierUploadPath); err != nil {
		hlog.CtxErrorf(ctx, "UploadFile failed, path: %v, err: %v", carrierUploadPath, err)
		utils.ResponseError(c, "UploadFile failed", err)
		return
	}
	// 上传数据文件
	dataUploadPath := filepath.Join(dir, "data.png")
	if err := fileService.UploadFile(ctx, c, DATA_FILE, dataUploadPath); err != nil {
		hlog.CtxErrorf(ctx, "UploadFile failed, path: %v, err: %v", dataUploadPath, err)
		utils.ResponseError(c, "UploadFile failed", err)
		return
	}
	// 获取加密文件
	resultFilePath := filepath.Join(dir, "result.png")
	err = stegService.EncodeImage(carrierUploadPath, dataUploadPath, resultFilePath)
	if err != nil {
		hlog.CtxErrorf(ctx, "EncodeImage failed, path: %v, err: %v", resultFilePath, err)
		utils.ResponseError(c, "EncodeImage failed", err)
		return
	}
	// 上传结果文件至Object Storage
	url, err := qiniu.PutFile(ctx, resultFilePath)
	if err != nil {
		hlog.CtxErrorf(ctx, "PutFile failed, path: %v, err: %v", resultFilePath, err)
		utils.ResponseError(c, "PutFile failed", err)
		return
	}
	// Response
	resp := map[string]interface{}{
		"carrier_file": carrierUploadPath,
		"data_file":    dataUploadPath,
		"result_file":  resultFilePath,
		"url":          url,
	}
	utils.ResponseOK(c, "EncodeImageFromImage Success", resp)
}

// DecodeImageFromImage 图片中解密图片
func DecodeImageFromImage(ctx context.Context, c *app.RequestContext) {
	// API限流
	remoteIp := utils.RemoteIp(c)
	hlog.CtxInfof(ctx, "Request: %v, remoteIp: %v", string(c.URI().Path()), remoteIp)
	if !redis.GetDecodeLock(ctx, remoteIp) {
		hlog.CtxErrorf(ctx, "GetDecodeLock failed, remoteIp: %v", remoteIp)
		utils.ResponseError(c, fmt.Sprintf("GetDecodeLock failed, remoteIp: %v", remoteIp), nil)
		return
	}
	// 创建临时文件夹
	key := "decode_image"
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
		_ = os.RemoveAll(path)
	}(dir)
	// 上传载体文件
	carrierUploadPath := filepath.Join(dir, "carrier.png")
	if err := fileService.UploadFile(ctx, c, CARRIER_FILE, carrierUploadPath); err != nil {
		hlog.CtxErrorf(ctx, "UploadFile failed, path: %v, err: %v", carrierUploadPath, err)
		utils.ResponseError(c, "UploadFile failed", err)
		return
	}
	// 获取被加密文件
	res, err := stegService.DecodeImage(carrierUploadPath)
	if err != nil {
		hlog.CtxErrorf(ctx, "DecodeImage failed, path: %v, err: %v", carrierUploadPath, err)
		utils.ResponseError(c, "DecodeImage failed", err)
		return
	}
	// 上传结果文件至Object Storage
	url, err := qiniu.PutFile(ctx, res)
	if err != nil {
		hlog.CtxErrorf(ctx, "PutFile failed, path: %v, err: %v", res, err)
		utils.ResponseError(c, "PutFile failed", err)
		return
	}
	// Response
	resp := map[string]interface{}{
		"carrier_file": carrierUploadPath,
		"result_file":  res,
		"url":          url,
	}
	utils.ResponseOK(c, "DecodeImageFromImage Success", resp)
}

// EncodeDocFromImage 图片中加密文字
func EncodeDocFromImage(ctx context.Context, c *app.RequestContext) {
	// API限流
	remoteIp := utils.RemoteIp(c)
	hlog.CtxInfof(ctx, "Request: %v, remoteIp: %v", string(c.URI().Path()), remoteIp)
	if !redis.GetEncodeLock(ctx, remoteIp) {
		hlog.CtxErrorf(ctx, "GetEncodeLock failed, remoteIp: %v", remoteIp)
		utils.ResponseError(c, fmt.Sprintf("GetEncodeLock failed, remoteIp: %v", remoteIp), nil)
		return
	}
	// 创建临时文件夹
	key := "encode_doc"
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
		_ = os.RemoveAll(path)
	}(dir)
	// 上传载体文件
	carrierUploadPath := filepath.Join(dir, "carrier.png")
	if err := fileService.UploadFile(ctx, c, CARRIER_FILE, carrierUploadPath); err != nil {
		hlog.CtxErrorf(ctx, "UploadFile failed, path: %v, err: %v", carrierUploadPath, err)
		utils.ResponseError(c, "UploadFile failed", err)
		return
	}
	// 获取文档数据
	dataDoc := c.PostForm(DATA_DOC)
	if dataDoc == "" {
		hlog.CtxErrorf(ctx, "data_doc is empty")
		utils.ResponseError(c, "data_doc is empty", nil)
		return
	}
	// 获取加密文件
	resultFilePath := filepath.Join(dir, "result.png")
	err = stegService.EncodeDoc(carrierUploadPath, dataDoc, resultFilePath)
	if err != nil {
		hlog.CtxErrorf(ctx, "EncodeDoc failed, path: %v, err: %v", resultFilePath, err)
		utils.ResponseError(c, "EncodeDoc failed", err)
		return
	}
	// 上传结果文件至Object Storage
	url, err := qiniu.PutFile(ctx, resultFilePath)
	if err != nil {
		hlog.CtxErrorf(ctx, "PutFile failed, path: %v, err: %v", resultFilePath, err)
		utils.ResponseError(c, "PutFile failed", err)
		return
	}
	// Response
	resp := map[string]interface{}{
		"carrier_file": carrierUploadPath,
		"data_doc":     dataDoc,
		"result_file":  resultFilePath,
		"url":          url,
	}
	utils.ResponseOK(c, "EncodeDocFromImage Success", resp)
}

// DecodeDocFromImage 图片中解密文字
func DecodeDocFromImage(ctx context.Context, c *app.RequestContext) {
	// API限流
	remoteIp := utils.RemoteIp(c)
	hlog.CtxInfof(ctx, "Request: %v, remoteIp: %v", string(c.URI().Path()), remoteIp)
	if !redis.GetDecodeLock(ctx, remoteIp) {
		hlog.CtxErrorf(ctx, "GetDecodeLock failed, remoteIp: %v", remoteIp)
		utils.ResponseError(c, fmt.Sprintf("GetDecodeLock failed, remoteIp: %v", remoteIp), nil)
		return
	}
	// 创建临时文件夹
	key := "decode_doc"
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
		_ = os.RemoveAll(path)
	}(dir)
	// 上传载体文件
	carrierUploadPath := filepath.Join(dir, "carrier.png")
	if err := fileService.UploadFile(ctx, c, CARRIER_FILE, carrierUploadPath); err != nil {
		hlog.CtxErrorf(ctx, "UploadFile failed, path: %v, err: %v", carrierUploadPath, err)
		utils.ResponseError(c, "UploadFile failed", err)
		return
	}
	// 获取被加密文档数据
	res, err := stegService.DecodeDoc(carrierUploadPath)
	if err != nil {
		hlog.CtxErrorf(ctx, "DecodeDoc failed, err: %v", err)
		utils.ResponseError(c, "DecodeDoc failed", err)
		return
	}
	// Response
	resp := map[string]interface{}{
		"carrier_file": carrierUploadPath,
		"result_doc":   res,
	}
	utils.ResponseOK(c, "DecodeDocFromImage Success", resp)
}

// DecodeDocOrImageFromImage 图片中智能解密图片或文字
func DecodeDocOrImageFromImage(ctx context.Context, c *app.RequestContext) {
	// API限流
	remoteIp := utils.RemoteIp(c)
	hlog.CtxInfof(ctx, "Request: %v, remoteIp: %v", string(c.URI().Path()), remoteIp)
	if !redis.GetDecodeLock(ctx, remoteIp) {
		hlog.CtxErrorf(ctx, "GetDecodeLock failed, remoteIp: %v", remoteIp)
		utils.ResponseError(c, fmt.Sprintf("GetDecodeLock failed, remoteIp: %v", remoteIp), nil)
		return
	}
	// 创建临时文件夹
	key := "decode_intelligent"
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
		_ = os.RemoveAll(path)
	}(dir)
	// 上传载体文件
	carrierUploadPath := filepath.Join(dir, "carrier.png")
	if err := fileService.UploadFile(ctx, c, CARRIER_FILE, carrierUploadPath); err != nil {
		hlog.CtxErrorf(ctx, "UploadFile failed, path: %v, err: %v", carrierUploadPath, err)
		utils.ResponseError(c, "UploadFile failed", err)
		return
	}
	// 获取被加密文件
	res, err := stegService.DecodeImage(carrierUploadPath)
	if err != nil {
		hlog.CtxErrorf(ctx, "DecodeImage failed, err: %v", err)
		// 获取被加密文档数据
		res, err = stegService.DecodeDoc(carrierUploadPath)
		if err != nil {
			hlog.CtxErrorf(ctx, "DecodeDoc failed, err: %v", err)
			utils.ResponseError(c, "DecodeDoc failed", errors.New("image中不存在图片或文字"))
			return
		}
		// Response
		resp := map[string]interface{}{
			"carrier_file": carrierUploadPath,
			"result_doc":   res,
		}
		utils.ResponseOK(c, "DecodeDocFromImage Success", resp)
		return
	}
	// 上传结果文件至Object Storage
	url, err := qiniu.PutFile(ctx, res)
	if err != nil {
		hlog.CtxErrorf(ctx, "PutFile failed, path: %v, err: %v", res, err)
		utils.ResponseError(c, "PutFile failed", err)
		return
	}
	// Response
	resp := map[string]interface{}{
		"carrier_file": carrierUploadPath,
		"result_file":  res,
		"url":          url,
	}
	utils.ResponseOK(c, "DecodeImageFromImage Success", resp)
}
