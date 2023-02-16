package connection

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/riicarus/loveshop/conf"
)

type RedisConnection[T interface{}] struct {
	Client *redis.Client
}

var RedisClient *redis.Client

func InitRedisConn() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        conf.ServiceConf.Cache.Redis.Addr,
		Password:    conf.ServiceConf.Cache.Redis.Password,
		DB:          conf.ServiceConf.Cache.Redis.DB,
		PoolSize:    conf.ServiceConf.Cache.Redis.PoolSize,
		DialTimeout: time.Millisecond * time.Duration(conf.ServiceConf.Cache.Redis.DialTimeout),
	})
}

func NewRedisConnection[T interface{}]() *RedisConnection[T] {
	return &RedisConnection[T]{
		Client: RedisClient,
	}
}

func (c *RedisConnection[T]) DoSimpleSet(key string, value T, expire int) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancel()

	return c.Client.Set(ctx, key, value, time.Duration(expire)).Err()
}

func (c *RedisConnection[T]) DoSimpleGet(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancel()

	return c.Client.Get(ctx, key).Result()
}

func (c *RedisConnection[T]) DoHashSet(k, hk string, value T, expire int) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancel()

	byteValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.Client.HSet(ctx, k, hk, string(byteValue)).Err()
}

// return err only when receiving non-redis.Nil err
func (c *RedisConnection[T]) DoHashGet(k, hk string, v T) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancel()

	cmder := c.Client.HGet(ctx, k, hk)

	if cmder.Err() == redis.Nil {
		return nil
	} else if cmder.Err() != nil {
		return cmder.Err()
	}

	err := json.Unmarshal([]byte(cmder.Val()), v)
	if err != nil {
		return err
	}

	return nil
}

func (c *RedisConnection[T]) DoHashMGet(k string, hks []string) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancel()

	cmder := c.Client.HMGet(ctx, k, hks...)

	return cmder.Val(), cmder.Err()
}

// temp should be a value of a struct, not a pointer
func (c *RedisConnection[T]) DoHashGetAll(k string, all []T) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancel()

	cmder := c.Client.HGetAll(ctx, k)

	if cmder.Err() == redis.Nil {
		return nil
	} else if cmder.Err() != nil {
		return cmder.Err()
	}

	t := new(T)
	for _, val := range cmder.Val() {
		err := json.Unmarshal([]byte(val), t)
		if err != nil {
			return err
		}

		all = append(all, *t)
	}

	return nil
}

func (c *RedisConnection[T]) DoHashGetAllMap(k string, all map[string]T) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancel()

	cmder := c.Client.HGetAll(ctx, k)

	if cmder.Err() == redis.Nil {
		return nil
	} else if cmder.Err() != nil {
		return cmder.Err()
	}

	t := new(T)
	for key, val := range cmder.Val() {
		err := json.Unmarshal([]byte(val), t)
		if err != nil {
			return err
		}

		all[key] = *t
	}

	return nil
}

func (c *RedisConnection[T]) DoHashRemove(k string, hk ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancel()

	cmder := c.Client.HDel(ctx, k, hk...)

	if cmder.Err() != nil {
		return cmder.Err()
	}

	return nil
}

func (c *RedisConnection[T]) DoAny(args []interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancel()

	return c.Client.Do(ctx, args...).Result()
}
