package tinify

import (
	"context"
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

// InitTinify 初始化tinify图片压缩client
func InitTinify() {
	// 配置初始化
	config = viper.Conf.Tinify

	var err error
	Client, err = client.NewClient()
	if err != nil {
		hlog.Fatalf("[Tinify] Init Tinify Failed, %v", err)
	}
	hlog.Info("[Tinify] Init Tinify Success")
}

func UploadImageZip(fbyte []byte) (shrinkResp *model.ShrinkResp, err error) {
	req := protocol.AcquireRequest()
	res := protocol.AcquireResponse()
	req.SetMethod(consts.MethodPost)
	req.SetRequestURI(config.Host + "/shrink")
	req.SetHeader("Authorization", config.Auth)
	req.SetBody(fbyte)
	if err = Client.Do(context.Background(), req, res); err != nil {
		return nil, err
	}
	shrinkResp = &model.ShrinkResp{}
	if err = json.Unmarshal(res.Body(), shrinkResp); err != nil {
		return nil, err
	}
	return shrinkResp, err
}
