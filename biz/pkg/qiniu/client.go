package qiniu

import (
	"context"
	"fmt"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/viper"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	formUploader *storage.FormUploader
	upToken      string
	config       *viper.Qiniu
)

// InitQiniu 初始化七牛云OS
func InitQiniu() {
	// 配置初始化
	config = viper.Conf.Qiniu

	// 初始化时刷新凭证
	RefreshToken()

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader = storage.NewFormUploader(&cfg)
	hlog.Info("Init Qiniu Success")
}

// PutFile 上传单文件
func PutFile(ctx context.Context, localFile string) (string, error) {
	ret := storage.PutRet{}
	hash, err := utils.GetFileHash(localFile)
	if err != nil {
		return "", err
	}
	err = formUploader.PutFile(ctx, &ret, upToken, fmt.Sprintf("%v/%v", config.Prefix, hash), localFile, nil)
	if err != nil {
		return "", err
	}
	hlog.CtxInfof(ctx, "Put File path:[%v], key:[%v], hash:[%v]", localFile, ret.Key, ret.Hash)
	return fmt.Sprintf("%v/%v", config.Domain, ret.Key), nil
}

// RefreshToken 刷新凭证
func RefreshToken() {
	putPolicy := storage.PutPolicy{
		Scope: config.Bucket,
	}
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)
	upToken = putPolicy.UploadToken(mac)
}

// GetUpToken 获取UpToken
func GetUpToken() string {
	return upToken
}
