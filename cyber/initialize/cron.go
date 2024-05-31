package initialize

import (
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/robfig/cron/v3"

	"github.com/DrReMain/cyber-base-server/cyber"
)

func Cron() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())

		{ // 模板定时任务
			_, err := cyber.Cron.AddTaskByFunc("background", "@every 1m", func() {
				hlog.Warnf("每一分钟打印一次日志 %s", time.Now().Format("2006-01-02 15:04:05"))
			}, "模板定时任务", option...)
			if err != nil {
				log.Fatalf("[Cron]: 添加定时任务失败: '%s'\n", err)
			}
		}

		//{
		//	_, err := cyber.Cron.AddTaskByFunc("定时任务名称", "cron表达式", func() {
		//		具体执行
		//	}, "任务名称", option...)
		//	if err != nil {
		//		log.Fatalf("[Cron]: 添加定时任务失败: '%s'\n", err)
		//	}
		//}

	}()
}
