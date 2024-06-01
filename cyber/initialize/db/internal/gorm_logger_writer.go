package internal

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm/logger"

	"github.com/DrReMain/cyber-base-server/cyber/config"
)

type Writer struct {
	config config.GeneralDB
	writer logger.Writer
}

func NewWriter(config config.GeneralDB, writer logger.Writer) *Writer {
	return &Writer{config: config, writer: writer}
}

func (c *Writer) Printf(message string, data ...any) {
	if c.config.LogZap {
		switch c.config.LogMode() {
		case logger.Silent:
		case logger.Error:
			hlog.Error(fmt.Sprintf(message, data...))
		case logger.Warn:
			hlog.Warn(fmt.Sprintf(message, data...))
		case logger.Info:
			fallthrough
		default:
			hlog.Info(fmt.Sprintf(message, data...))
		}
		return
	}
	c.writer.Printf(message, data...)
}
