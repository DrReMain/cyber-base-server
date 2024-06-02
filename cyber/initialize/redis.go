package initialize

import (
	"context"
	"log"

	"github.com/DrReMain/cyber-base-server/cyber"
	"github.com/redis/go-redis/v9"
)

func Redis() {
	if cyber.Config.System.Redis {
		redisCfg := cyber.Config.Redis
		var client redis.UniversalClient

		if redisCfg.Cluster {
			client = redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:    redisCfg.ClusterAddress,
				Password: redisCfg.Password,
			})
		} else {
			client = redis.NewClient(&redis.Options{
				Addr:     redisCfg.Address,
				Password: redisCfg.Password,
				DB:       redisCfg.DB,
			})
		}

		pong, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("[Redis]: 连接Redis失败 -> '%s'\n", err)
		} else {
			log.Printf("[Redis]: ping response -> %s\n", pong)
			cyber.Redis = client
		}
	}
}
