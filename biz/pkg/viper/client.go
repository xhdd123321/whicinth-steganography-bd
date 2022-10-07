package viper

import (
	"context"
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
)

var Conf *Config

// InitViper 初始化Viper
func InitViper() {
	ctx := context.Background()
	viper.SetConfigType("yaml")
	runEnv := os.Getenv("RUN_ENV")
	if runEnv == "DEV" {
		viper.SetConfigFile("./conf/dev.config.yaml")
	} else if runEnv == "PROD" {
		viper.SetConfigFile("./conf/prod.config.yaml")
	} else {
		viper.SetConfigFile("./conf/default.config.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		hlog.CtxErrorf(ctx, "[Viper]ReadInConfig failed, err: %v", err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		hlog.CtxErrorf(ctx, "[Viper]Unmarshal failed, err: %v", err)
	}

	hlog.CtxInfof(ctx, "Conf.App: %#v", Conf.App)
	hlog.CtxInfof(ctx, "Conf.Redis: %#v", Conf.Redis)
	hlog.CtxInfof(ctx, "Conf.Qiniu: %#v", Conf.Qiniu)
	hlog.CtxInfof(ctx, "Conf.Cronjob: %#v", Conf.Cronjob)
}
