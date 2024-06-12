package main

import (
	"fmt"
	"log"
	"math"

	"github.com/gomodule/redigo/redis"
)

type RedisPool interface {
	Get() redis.Conn
}

type Cache struct {
	// redisPool *redis.Pool
	redisPool RedisPool
	enabled   bool
}

func (c Cache) InitPool(redisHost, redisPort string) RedisPool {
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	const maxConnections = 10

	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", redisAddr)
	}, maxConnections)

	return pool
}

func main() {
	log.Println(math.Sqrt(64))
}
