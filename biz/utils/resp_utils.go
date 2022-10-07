package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func ResponseOK(c *app.RequestContext, msg string, data interface{}) {
	c.JSON(consts.StatusOK, utils.H{
		"code":    0,
		"message": msg,
		"data":    data,
	})
}

func ResponseError(c *app.RequestContext, msg string, err error) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	c.JSON(consts.StatusInternalServerError, utils.H{
		"code":    1,
		"message": msg,
		"data": utils.H{
			"err_msg": errMsg,
		},
	})
}
