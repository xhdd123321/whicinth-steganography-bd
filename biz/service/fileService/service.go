package fileService

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
)

func UploadFile(ctx context.Context, c *app.RequestContext, fileField string, uploadPath string) error {
	file, err := c.FormFile(fileField)
	if err != nil {
		logger.CtxErrorf(ctx, "Get FormFile failed, err: %v", err)
		return err
	}
	err = c.SaveUploadedFile(file, uploadPath)
	if err != nil {
		logger.CtxErrorf(ctx, "SaveUploadedFile failed, filename: %s, err: %v", file.Filename, err)
		return err
	}
	return nil
}
