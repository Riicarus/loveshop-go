package connection

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/riicarus/loveshop/conf"
)

type RedisConnction[T interface{}] struct {
	Client *redis.Client
}

var client *redis.Client

func InitRedisConn() {
	client = redis.NewClient(&redis.Options{
		Addr:        conf.ServiceConf.Cache.Redis.Addr,
		Password:    conf.ServiceConf.Cache.Redis.Password,
		DB:          conf.ServiceConf.Cache.Redis.DB,
		PoolSize:    conf.ServiceConf.Cache.Redis.PoolSize,
		DialTimeout: time.Millisecond * time.Duration(conf.ServiceConf.Cache.Redis.DialTimeout),
	})
}

func NewRedisConnection[T interface{}]() *RedisConnction[T] {
	return &RedisConnction[T]{
		Client: client,
	}
}

func (c *RedisConnction[T]) DoSimpleSet(key string, value T, expire int) error {
	ctx, cancle := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancle()

	return c.Client.Set(ctx, key, value, time.Duration(expire)).Err()
}

func (c *RedisConnction[T]) DoSimpleGet(key string) (string, error) {
	ctx, cancle := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancle()

	return c.Client.Get(ctx, key).Result()
}

func (c *RedisConnction[T]) DoHashSet(k, hk string, value T, expire int) error {
	ctx, cancle := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancle()

	byteValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.Client.HSet(ctx, k, hk, string(byteValue)).Err()
}

// return err only when receiving non-redis.Nil err
func (c *RedisConnction[T]) DoHashGet(k, hk string, v T) error {
	ctx, cancle := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancle()

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

// temp should be a value of a struct, not a pointer
func (c *RedisConnction[T]) DoHashGetAll(k string, temp T, all []T) error {
	ctx, cancle := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancle()

	cmder := c.Client.HGetAll(ctx, k)

	if cmder.Err() == redis.Nil {
		return nil
	} else if cmder.Err() != nil {
		return cmder.Err()
	}

	for _, val := range cmder.Val() {
		err := json.Unmarshal([]byte(val), &temp)
		if err != nil {
			return err
		}

		all = append(all, temp)
	}

	return nil
}

func (c *RedisConnction[T]) DoHashRemove(k string, hk ...string) error {
	ctx, cancle := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancle()

	cmder := c.Client.HDel(ctx, k, hk...)

	if cmder.Err() != nil {
		return cmder.Err()
	}

	return nil
}

func (c *RedisConnction[T]) DoAny(args []interface{}) (interface{}, error) {
	ctx, cancle := context.WithTimeout(context.Background(), c.Client.Options().DialTimeout)
	defer cancle()

	return c.Client.Do(ctx, args...).Result()
}