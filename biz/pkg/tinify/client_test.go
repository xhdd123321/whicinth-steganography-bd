package tinify

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/viper"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("RUN_ENV", "DEV")
	viper.InitViper()
	InitTinify()
	m.Run()
}

func TestUploadImageZip(t *testing.T) {
	testFilePath := filepath.Join(utils.GetMediaAbPath(), "test", "data_req.png")
	fByte, err := os.ReadFile(testFilePath)
	if err != nil {
		t.Errorf("ReadFile failed, err: %v", err)
	}
	res, err := UploadImage2Compare(context.Background(), fByte)
	if err != nil {
		t.Errorf("ReadFile failed, err: %v", err)
	}
	t.Logf("upload image zip success: %v", res.String())
}
