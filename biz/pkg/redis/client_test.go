package redis

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/viper"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("RUN_ENV", "DEV")
	viper.InitViper()
	InitRedis()
	m.Run()
}

func TestGetIncrId(t *testing.T) {
	res := GetIncrId(context.Background(), "order")
	assert.NotEqual(t, 0, res)
	fmt.Println("get id success: ", res)
}
