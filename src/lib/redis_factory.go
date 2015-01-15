package lib

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"runtime"
)

func NewFactory(host string) *RedisConf {
	v := &RedisConf{
		Host: host,
	}
	v.Pool = v.NewPool()

	return v
}

func (conf *RedisConf) NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, //Maximum number of connections
		Dial: func() (redis.Conn, error) {
			if c, err := redis.Dial("tcp", conf.Host); err != nil {
				panic(err.Error())
			} else {
				return c, nil
			}
		},
	}
}

func (conf *RedisConf) Add(item *ShortenedURL) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	c := conf.Pool.Get()
	defer c.Close()

	fmt.Printf("s: %v\n", item)

	c.Send("MULTI")
	if s, err := json.Marshal(item); err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("s: %s\n", string(s))
		c.Send("SET", item.Key, s)

		if item.Expires > 0 {
			c.Send("EXPIRE", item.Key, item.Expires)
		}
		if _, err := c.Do("EXEC"); err != nil {
			panic(err.Error())
		}
	}

	return nil
}

func (conf *RedisConf) Get(key string, e interface{}) (interface{}, error) {
	c := conf.Pool.Get()
	defer c.Close()

	if r, err := c.Do("GET", key); err != nil {
		return nil, err
	} else {
		if err = json.Unmarshal(r.([]byte), e); err != nil {
			return nil, err
		}

		return e, nil
	}
}

func (conf *RedisConf) Exists(key string) (int64, error) {
	c := conf.Pool.Get()
	defer c.Close()

	if r, err := c.Do("EXISTS", key); err != nil {
		return 0, err
	} else {
		return r.(int64), nil
	}
}

func (conf *RedisConf) RPush(key string, val ...string) (int64, error) {
	c := conf.Pool.Get()
	defer c.Close()

	if r, err := c.Do("LPUSH", key, val); err != nil {
		return -1, err
	} else {
		return r.(int64), nil
	}
}
