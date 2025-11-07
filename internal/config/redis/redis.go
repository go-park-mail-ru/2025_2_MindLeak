package redis

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

type RedisConfig struct {
	redisAddr string
	redisPass string
	pool      *redis.Pool
}

func NewRedisConfig() *RedisConfig {
	addr := os.Getenv("REDIS_ADDR")
	pass := os.Getenv("REDIS_PASSWORD")

	cfg := &RedisConfig{
		redisAddr: addr,
		redisPass: pass,
	}

	cfg.pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialPassword(pass))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	logger.Info(nil, "✅ Redis pool initialized at %s", addr)
	return cfg
}

func (r *RedisConfig) GetConn() redis.Conn {
	return r.pool.Get()
}

func (r *RedisConfig) Ping() error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("PING")
	return err
}

func (r *RedisConfig) GetPool() *redis.Pool {
	return r.pool
}
