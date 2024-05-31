package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/DrReMain/cyber-base-server/cyber"
	"github.com/DrReMain/cyber-base-server/cyber/initialize/db/internal"
)

func GormPgsql() *gorm.DB {
	p := cyber.Config.Pgsql
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}
