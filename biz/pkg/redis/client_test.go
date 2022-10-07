package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Redis Test Start")
	InitRedis()
	m.Run()
}

func TestGetOrderId(t *testing.T) {
	res := GetIncrId(context.Background(), "order")
	assert.NotEqual(t, 0, res)
	fmt.Println("get id success: ", res)
}
