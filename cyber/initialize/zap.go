package initialize

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hzzap "github.com/hertz-contrib/logger/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/DrReMain/cyber-base-server/cyber"
)

func Zap() {
	if err := os.MkdirAll(cyber.Config.Zap.Director, os.ModePerm); err != nil {
		log.Fatalf("[Zap]: 创建日志目录失败 '%s'\n", err)
	}

	var opts []zap.Option
	if cyber.Config.Zap.ShowLine {
		opts = []zap.Option{
			zap.AddCaller(),
			zap.AddCallerSkip(3),
		}
	}

	logger := hzzap.NewLogger(
		hzzap.WithCores(buildCores()...),
		hzzap.WithZapOptions(opts...),
	)

	cyber.Logger = logger
	hlog.SetLogger(logger)
}

func buildCores() []hzzap.CoreConfig {
	levels := cyber.Config.Zap.Levels()
	length := len(levels)
	cores := make([]hzzap.CoreConfig, 0, length)
	for i := 0; i < length; i++ {
		cores = append(cores, buildCore(levels[i]))
	}
	return cores
}

func buildCore(level zapcore.Level) hzzap.CoreConfig {
	return hzzap.CoreConfig{
		Enc: cyber.Config.Zap.Encoder(),
		Ws:  buildWriteSyncer(level),
		Lvl: zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == level
		}),
	}
}

func buildWriteSyncer(level zapcore.Level) zapcore.WriteSyncer {
	c := buildCutter(level.String())
	if cyber.Config.Zap.Stdout {
		multi := zapcore.NewMultiWriteSyncer(os.Stdout, c)
		return zapcore.AddSync(multi)
	}
	return zapcore.AddSync(c)
}

type cutter struct {
	level  string   // debug, info, warn, error, dpanic, panic, fatal
	format string   // 时间格式 2006-01-02 15:04:05
	dir    string   // 日志文件夹
	file   *os.File // 文件句柄
	mutex  *sync.RWMutex
}

func buildCutter(level string) *cutter {
	return &cutter{
		level:  level,
		format: time.DateOnly,
		dir:    cyber.Config.Zap.Director,
		mutex:  new(sync.RWMutex),
	}
}

func (c *cutter) Write(bytes []byte) (n int, err error) {
	c.mutex.Lock()
	defer func() {
		if c.file != nil {
			_ = c.file.Close()
			c.file = nil
		}
		c.mutex.Unlock()
	}()

	filename := filepath.Join(c.dir, time.Now().Format(c.format)+"."+c.level+".log")
	director := filepath.Dir(filename)
	err = os.MkdirAll(director, os.ModePerm)
	if err != nil {
		return 0, err
	}
	c.file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return 0, err
	}
	return c.file.Write(bytes)
}

func (c *cutter) Sync() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.file != nil {
		return c.file.Sync()
	}
	return nil
}
