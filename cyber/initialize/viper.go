package initialize

import (
	"flag"
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/DrReMain/cyber-base-server/cyber"
)

var (
	configFileTest    = "config.test.yaml"
	configFileDebug   = "config.debug.yaml"
	configFileRelease = "config.release.yaml"
)

func Viper() {
	var configFile string

	flag.StringVar(&configFile, "c", "", "指定配置文件")
	flag.Parse()

	if configFile == "" {
		if env := os.Getenv("CYBER_ENV"); env == "" {
			configFile = configFileDebug
		} else {
			switch env {
			case "test":
				configFile = configFileTest
			case "debug":
				configFile = configFileDebug
			case "release":
				configFile = configFileRelease
			}
		}
	}
	configFile = "etc/" + configFile
	log.Printf("[Viper]: 正在使用'%s'\n", configFile)

	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("[Viper]: 读取配置文件失败: %s\n", err)
	}
	v.WatchConfig()
	if err = v.Unmarshal(&cyber.Config); err != nil {
		log.Fatalf("[Viper]: 反序列化配置文件失败: %s\n", err)
	}

	cyber.Viper = v
}
