package middleware

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	cutils_jwt "github.com/DrReMain/cyber-base-server/cyber/utils/jwt"
)

func JsonWebToken() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		accessToken := string(c.GetHeader("Authorization"))
		j := cutils_jwt.NewJsonWebToken()
		_, err := j.ParseToken(accessToken)
		if err != nil {
			var status int
			var code string
			switch {
			case errors.Is(err, cutils_jwt.TokenBuffer):
				status = 498
				code = "200001"
			default:
				status = consts.StatusUnauthorized
				code = "200000"
			}
			c.JSON(status, utils.H{
				"base": utils.H{
					"t":       time.Now().UnixMilli(),
					"success": false,
					"code":    code,
					"msg":     err.Error(),
				},
				"result": nil,
			})
			c.Abort()
			return
		}

		c.Next(ctx)
	}
}
