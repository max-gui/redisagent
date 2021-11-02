package redisops

import (
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

func Pool() *redis.Pool {
	once.Do(func() {

		pool = &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
			Dial: func() (redis.Conn, error) { return redis.Dial("tcp", url, redis.DialPassword(pwd)) },
		}
	})

	return pool
}

var (
	pool     *redis.Pool
	once     sync.Once
	url, pwd string
	// redisServer *string
)

// func init() {
// 	// testing.Init()
// 	// redisServer = flag.String("redisServer", ":6379", "") //"10.25.80.6:6480", "") //":6379", "")
// 	// flag.Parse()
// 	Pool = NewPool(*constset.RedisServer)
// }
