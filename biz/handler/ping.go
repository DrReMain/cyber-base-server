// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Ping .
func Ping(ctx context.Context, c *app.RequestContext) {
	hlog.Info("Ping Success!!!")
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}
