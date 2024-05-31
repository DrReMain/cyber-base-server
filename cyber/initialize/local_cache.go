package initialize

import (
	"log"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/DrReMain/cyber-base-server/cyber"
	"github.com/DrReMain/cyber-base-server/cyber/utils/h_duration"
)

func LocalCache() {
	du, err := cutils_hd.ParseDuration(cyber.Config.Jwt.ExpiresTime)
	if err != nil {
		log.Fatalf("[Utils]: 转换Jwt.ExpiresTime失败 '%s'\n", err)
	}

	cyber.LocalCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(du))
}
