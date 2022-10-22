package tinify

import (
	"context"
	"crypto/tls"
	"encoding/json"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/model"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/viper"
)

var (
	Client *client.Client
	config *viper.Tinify
)

// InitTinify 初始化Tinify图片压缩Client
func InitTinify() {
	// 配置初始化
	config = viper.Conf.Tinify

	var err error
	clientCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	Client, err = client.NewClient(
		client.WithTLSConfig(clientCfg),
	)
	if err != nil {
		hlog.Fatalf("[Tinify] Init Tinify Failed, %v", err)
	}
	hlog.Info("[Tinify] Init Tinify Success, host: %v", config.Host)
}

func UploadImage2Compare(ctx context.Context, fByte []byte) (shrinkResp *model.ShrinkResp, err error) {
	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(res)
	}()
	req.SetMethod(consts.MethodPost)
	req.SetRequestURI(config.Host + "/shrink")
	req.SetHeader("Authorization", config.Auth)
	req.SetBody(fByte)
	if err = Client.Do(context.Background(), req, res); err != nil {
		hlog.CtxErrorf(ctx, "[Tinify] Request shrink API, err: %v", err)
		return nil, err
	}
	shrinkResp = &model.ShrinkResp{}
	if err = json.Unmarshal(res.Body(), shrinkResp); err != nil {
		return nil, err
	}
	return shrinkResp, err
}
