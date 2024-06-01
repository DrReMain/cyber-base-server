//go:build migrate

package main

import (
	"log"

	"github.com/DrReMain/cyber-base-server/cyber"
	"github.com/DrReMain/cyber-base-server/cyber/initialize"

	// gorm models
	"github.com/DrReMain/cyber-base-server/biz/dal/sys_model"
)

func init() {
	initialize.Viper()
	initialize.Gorm()
}

func main() {
	if cyber.DB != nil {
		db, _ := cyber.DB.DB()
		defer db.Close()

		cyber.DB.AutoMigrate(
			sys_model.SysUser{},
			sys_model.SysDept{},
		)
	} else {
		log.Fatalln("[DB]: 未连接数据库")
	}
}
