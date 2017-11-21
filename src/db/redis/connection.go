package redis

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

var (
	// 定义常量
	RedisClient    *redis.Pool
	REDIS_HOST     string
	REDIS_DB       int
	REDIS_PASSWORD string
)

func init() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = beego.AppConfig.String("redis.host")
	REDIS_DB, _ = beego.AppConfig.Int("redis.db")
	REDIS_PASSWORD = beego.AppConfig.String("redis.password")
	logs.Debug("redis host, db", REDIS_HOST, REDIS_DB)
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 20),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 100),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			var c, err = redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// 选择db
			if REDIS_PASSWORD != "" {
				if _, err := c.Do("AUTH", REDIS_PASSWORD); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", REDIS_DB); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}
