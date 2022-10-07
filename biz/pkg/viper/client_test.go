package viper

import (
	"os"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("RUN_ENV", "TEST")
	m.Run()
}

func TestInitViper(t *testing.T) {
	InitViper()
	assert.NotNil(t, Conf)
}
