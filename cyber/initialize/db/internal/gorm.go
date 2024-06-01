package internal

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/DrReMain/cyber-base-server/cyber"
	"github.com/DrReMain/cyber-base-server/cyber/config"
)

var Gorm = new(_gorm)

type _gorm struct{}

func (g *_gorm) Config() *gorm.Config {
	var general config.GeneralDB
	switch cyber.Config.System.DB {
	case "mysql":
		general = cyber.Config.Mysql.GeneralDB
	case "pgsql":
		general = cyber.Config.Pgsql.GeneralDB
	case "sqlite":
		general = cyber.Config.Sqlite.GeneralDB
	default:
		log.Fatalln("[DB]: 未配置数据库类型")
	}
	return &gorm.Config{
		Logger: logger.New(
			NewWriter(general, log.New(os.Stdout, "\r\n", log.LstdFlags)),
			logger.Config{
				SlowThreshold: 200 * time.Millisecond,
				LogLevel:      general.LogMode(),
				Colorful:      true,
			}),
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
