package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

// RedisConf redis conf
type RedisConf struct {
	redis.Options
}

// Conn redis conn
type Conn struct {
	*redis.Client
}

// Dial dial redis and return conn
func Dial(conf *RedisConf) (*Conn, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, errors.WithMessage(err, "dial to redis error")
	}
	return &Conn{client}, nil
}
