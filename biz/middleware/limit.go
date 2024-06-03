package middleware

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/DrReMain/cyber-base-server/cyber"
)

type Limit struct {
	count  int
	expire time.Duration
}

func NewLimit() *Limit {
	return &Limit{
		count:  cyber.Config.System.IpLimitCount,
		expire: time.Duration(cyber.Config.System.IpLimitTime) * time.Second,
	}
}

func (l *Limit) HandlerFunc(ctx context.Context, c *app.RequestContext) {
	if err := l.setLimit(
		ctx,
		fmt.Sprintf("LIMIT(%s)", c.ClientIP()),
		l.count,
		l.expire,
	); err != nil {
		c.JSON(consts.StatusTooManyRequests, utils.H{
			"base": utils.H{
				"t":       time.Now().UnixMilli(),
				"success": false,
				"code":    "200002",
				"msg":     err.Error(),
			},
			"result": nil,
		})
		c.Abort()
		return
	} else {
		c.Next(ctx)
	}
}

func (l *Limit) setLimit(ctx context.Context, key string, count int, expire time.Duration) error {
	if cyber.Redis == nil {
		return nil
	}

	unity := errors.New("缓存服务异常，请联系管理员")

	c, err := cyber.Redis.Exists(ctx, key).Result()
	if err != nil {
		return unity
	}

	if c == 0 {
		pipe := cyber.Redis.TxPipeline()
		pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, expire)
		_, err = pipe.Exec(ctx)
		if err != nil {
			return unity
		}
		return nil
	}

	times, err := cyber.Redis.Get(ctx, key).Int()
	if err != nil {
		return unity
	}

	if times >= count {
		return errors.New("请求过于频繁，请稍后再试")
	}

	if cyber.Redis.Incr(ctx, key).Err() != nil {
		return unity
	}
	return nil
}
