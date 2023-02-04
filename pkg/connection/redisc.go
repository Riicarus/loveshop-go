package connection

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/riicarus/loveshop/conf"
)

type RedisConnction struct {
	Client *redis.Client
}

var RedisConn *RedisConnction

func InitRedisConn() {
	RedisConn = &RedisConnction{
		Client: redis.NewClient(&redis.Options{
			Addr:        conf.ServiceConf.Cache.Redis.Addr,
			Password:    conf.ServiceConf.Cache.Redis.Password,
			DB:          conf.ServiceConf.Cache.Redis.DB,
			PoolSize:    conf.ServiceConf.Cache.Redis.PoolSize,
			DialTimeout: time.Millisecond * time.Duration(conf.ServiceConf.Cache.Redis.DialTimeout),
		}),
	}
}

func (c *RedisConnction) DoSimpleSet(key string, value interface{}, expire int) error {
	ctx, cancle := context.WithTimeout(context.Background(), RedisConn.Client.Options().DialTimeout)
	defer cancle()

	return RedisConn.Client.Set(ctx, key, value, time.Duration(expire)).Err()
}

func (c *RedisConnction) DoSimpleGet(key string) (string, error) {
	ctx, cancle := context.WithTimeout(context.Background(), RedisConn.Client.Options().DialTimeout)
	defer cancle()

	return RedisConn.Client.Get(ctx, key).Result()
}

func (c *RedisConnction) DoHashSet(k, hk string, value interface{}, expire int) error {
	ctx, cancle := context.WithTimeout(context.Background(), RedisConn.Client.Options().DialTimeout)
	defer cancle()

	byteValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return RedisConn.Client.HSet(ctx, k, hk, string(byteValue)).Err()
}

// return err only when receiving non-redis.Nil err
func (c *RedisConnction) DoHashGet(k, hk string, v interface{}) error {
	ctx, cancle := context.WithTimeout(context.Background(), RedisConn.Client.Options().DialTimeout)
	defer cancle()

	cmder := RedisConn.Client.HGet(ctx, k, hk)

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
func (c *RedisConnction) DoHashGetAll(k string, temp interface{}, all []interface{}) error {
	ctx, cancle := context.WithTimeout(context.Background(), RedisConn.Client.Options().DialTimeout)
	defer cancle()

	cmder := RedisConn.Client.HGetAll(ctx, k)

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

func (c *RedisConnction) DoHashRemove(k string, hk ...string) error {
	ctx, cancle := context.WithTimeout(context.Background(), RedisConn.Client.Options().DialTimeout)
	defer cancle()

	cmder := RedisConn.Client.HDel(ctx, k, hk...)

	if cmder.Err() != nil {
		return cmder.Err()
	}

	return nil
}

func (c *RedisConnction) DoAny(args []interface{}) (interface{}, error) {
	ctx, cancle := context.WithTimeout(context.Background(), RedisConn.Client.Options().DialTimeout)
	defer cancle()

	return RedisConn.Client.Do(ctx, args...).Result()
}