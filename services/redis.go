/**
 * @author: hqd
 * @description: redis
 * @file: redis
 * @date: 2021-02-05 11:54
 */

package services

import (
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/garyburd/redigo/redis"
)

var RedisPool *redis.Pool

func InitRedis() error {
	redisAddr, _ := web.AppConfig.String("redisaddr")
	redisPwd, _ := web.AppConfig.String("redispwd")
	redisMaxIdle, _ := web.AppConfig.Int("redismaxidle")
	redisMaxActive, _ := web.AppConfig.Int("redismaxactive")

	RedisPool = &redis.Pool{
		MaxActive:       redisMaxActive, // 最大连接数
		MaxIdle:         redisMaxIdle,   // 最大的空闲连接数
		IdleTimeout:     time.Second * 2,
		Wait:            true,
		MaxConnLifetime: 0,
		Dial: func() (redis.Conn, error) {
			password := redis.DialPassword(redisPwd)
			timeout := redis.DialConnectTimeout(time.Second * 5)
			return redis.Dial("tcp", redisAddr, password, timeout)
		},
	}
	return poolTestConn(RedisPool)
}

func poolTestConn(pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		return err
	}
	logs.Info("init redis connection success")
	return nil
}
