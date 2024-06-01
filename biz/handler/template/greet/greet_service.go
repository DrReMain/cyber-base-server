// Code generated by hertz generator.

package greet

import (
	"context"

	"github.com/DrReMain/cyber-base-server/biz/common/res"
	greet "github.com/DrReMain/cyber-base-server/biz/hertz_gen/template/greet"

	"github.com/cloudwego/hertz/pkg/app"
)

// Greet .
// @router /v1/template/greet/:name_content [GET]
func Greet(ctx context.Context, c *app.RequestContext) {
	var err error
	var req greet.GreetReq
	err = c.BindAndValidate(&req)
	if err != nil {
		o := &greet.GreetRes{
			Base:   res.BaseValidateFail(err),
			Result: nil,
		}
		res.ValidateFail(c, o, err, res.Json(req))
		return
	}

	res.Success(c, &greet.GreetRes{
		Base:   res.BaseSuccess(),
		Result: &greet.Result{TextContent: "hello " + req.NameContent},
	})
}
