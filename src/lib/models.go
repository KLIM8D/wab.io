package lib

import (
	"github.com/garyburd/redigo/redis"
)

type ShortenedURL struct {
	Key     string
	Expires float64
	Url     string
}

type RedisConf struct {
	Pool *redis.Pool
	Host string
}

type Stats struct {
	Enabled           bool
	TotalUrlsKey      string
	TotalRedirectsKey string
}

type WebConf struct {
	Base string
	Port uint16
}

type Configuration struct {
	Redis RedisConf
	Stats Stats
	Web   WebConf
}

type User struct {
	Key     string
	Expires float64
}
