// Code generated by hertz generator.
//go:build !migrate

package main

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/DrReMain/cyber-base-server/cyber"
	"github.com/DrReMain/cyber-base-server/cyber/initialize"
)

func init() {
	initialize.Viper()
	initialize.Zap()
	initialize.LocalCache()
	initialize.Gorm()
	initialize.Redis()
	initialize.Cron()
}

func main() {
	if cyber.Logger != nil {
		defer cyber.Logger.Sync()
	}
	if cyber.DB != nil {
		db, _ := cyber.DB.DB()
		defer db.Close()
	}

	h := server.Default(
		server.WithHostPorts(fmt.Sprintf(":%d", cyber.Config.System.Port)),
	)

	register(h)
	h.Spin()
}
