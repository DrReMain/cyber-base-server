package cyber

import (
	"sync"

	hzzap "github.com/hertz-contrib/logger/zap"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/DrReMain/cyber-base-server/cyber/config"
	"github.com/DrReMain/cyber-base-server/cyber/utils/cron_task"
)

var (
	lock sync.RWMutex

	Config     config.Config
	Viper      *viper.Viper
	Logger     *hzzap.Logger
	LocalCache local_cache.Cache
	Redis      redis.UniversalClient
	DB         *gorm.DB

	Cron cutils_cron.Cron = cutils_cron.NewCronTask()
)
