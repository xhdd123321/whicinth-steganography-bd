package fileService

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/model"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/tinify"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"
)

// UploadFile 上传文件至本地
func UploadFile(ctx context.Context, c *app.RequestContext, fileField string, uploadPath string) error {
	file, err := c.FormFile(fileField)
	if err != nil {
		hlog.CtxErrorf(ctx, "Get FormFile failed, err: %v", err)
		return err
	}
	err = c.SaveUploadedFile(file, uploadPath)
	if err != nil {
		hlog.CtxErrorf(ctx, "SaveUploadedFile failed, filename: %s, err: %v", file.Filename, err)
		return err
	}
	return nil
}

// ClearFile 清理过期文件
func ClearFile(ctx context.Context, ttlMinutes float64) error {
	var needRemovePathList []string
	if err := filepath.WalkDir(utils.GetMediaAbPath(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && utils.IsNum(d.Name()) {
			if dirInfo, err := d.Info(); err == nil {
				if time.Since(dirInfo.ModTime()).Minutes() > ttlMinutes {
					needRemovePathList = append(needRemovePathList, path)
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}
	var success int
	for _, path := range needRemovePathList {
		if err := os.RemoveAll(path); err != nil {
			hlog.CtxErrorf(ctx, "Remove failed, path: %v, err: %v", path, err)
		} else {
			success++
			hlog.CtxInfof(ctx, "Remove success, path: %v", path)
		}
	}
	hlog.CtxInfof(ctx, "ClearFile Result: %v/%v", success, len(needRemovePathList))
	return nil
}

// CompressImage 无损压缩图片文件
func CompressImage(ctx context.Context, preFile string) (shrinkResp *model.ShrinkResp, err error) {
	fByte, err := os.ReadFile(preFile)
	if err != nil {
		return nil, fmt.Errorf("read preFile [%s] failed, err: %v", preFile, err)
	}
	res, err := tinify.UploadImage2Compare(ctx, fByte)
	if err != nil {
		hlog.CtxErrorf(ctx, "UploadImage2Compare failed, err: %v", err)
		return nil, err
	}
	return res, nil
}
