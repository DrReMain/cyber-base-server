// Code generated by hertz generator.

package greet

import (
	"context"

	greet "github.com/DrReMain/cyber-base-server/biz/hertz_gen/template/greet"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Greet .
// @router /v1/template/greet/:name_content [GET]
func Greet(ctx context.Context, c *app.RequestContext) {
	var err error
	var req greet.GreetReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(greet.GreetRes)

	c.JSON(consts.StatusOK, resp)
}