package redis

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/gomodule/redigo/redis"
	"os"
)

type RedisConfig struct {
	redisAddr string
	redisPass string
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		redisAddr: os.Getenv("REDIS_ADDR"),
		redisPass: os.Getenv("REDIS_PASSWORD"),
	}
}

func (conn *RedisConfig) RedisConnect() (redis.Conn, error) {
	c, err := redis.Dial(
		"tcp",
		conn.redisAddr,
		redis.DialPassword(conn.redisPass),
	)
	if err != nil {
		logger.Error(nil, "Error connecting to Redis: %v", err)
		return nil, err
	}

	_, err = c.Do("PING")
	if err != nil {
		logger.Error(nil, "Error pinging Redis: %v", err)
		_ = c.Close()
		return nil, err
	}

	logger.Info(nil, "✅ Connected to Redis at %s", conn.redisAddr)
	return c, nil
}

func (p *RedisConfig) GetAddr() string {
	return p.redisAddr
}

func (p *RedisConfig) GetPass() string {
	return p.redisPass
}
