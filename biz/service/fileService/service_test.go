package fileService

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"
)

const TTL_MINUTE_TEST = 1

func TestClearFile(t *testing.T) {
	err := ClearFile(context.Background(), TTL_MINUTE_TEST)
	if err != nil {
		fmt.Printf("ClearFile failed, err: %v", err)
		assert.NoError(t, err)
	}
	err = filepath.WalkDir(utils.GetMediaAbPath(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && utils.IsNum(d.Name()) {
			if dirInfo, err := d.Info(); err == nil {
				if time.Since(dirInfo.ModTime()).Minutes() > TTL_MINUTE_TEST {
					return errors.New("存在未删除过期文件")
				}
			}
		}
		return nil
	})
	assert.NoError(t, err)
}
