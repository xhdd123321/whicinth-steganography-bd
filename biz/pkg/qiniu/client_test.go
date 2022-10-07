package qiniu

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/viper"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Qiniu Test Start")
	_ = os.Setenv("RUN_ENV", "DEV")
	viper.InitViper()
	InitQiniu()
	m.Run()
}

func TestPutFile(t *testing.T) {
	res, err := PutFile(context.Background(), filepath.Join(utils.GetMediaAbPath(), "test", "image_with_doc.png"))
	assert.NoError(t, err)
	fmt.Println("put file success: ", res)
}
