package initialize

import (
	"log"

	"github.com/DrReMain/cyber-base-server/cyber"
	"github.com/DrReMain/cyber-base-server/cyber/initialize/db"
)

func Gorm() {
	switch cyber.Config.System.DB {
	case "mysql":
		cyber.DB = db.GormMysql()
	case "pgsql":
		cyber.DB = db.GormPgsql()
	case "sqlite":
		cyber.DB = db.GormSqlite()
	default:
		log.Fatalln("[DB]: 未配置数据库类型")
	}
}
