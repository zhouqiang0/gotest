package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

//定义一个全局pool
var pool *redis.Pool

func initPool(addr string, maxIdle int, maxActice int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		TestOnBorrow:    nil,
		MaxIdle:         maxIdle,
		MaxActive:       maxActice,
		IdleTimeout:     idleTimeout,
		Wait:            false,
		MaxConnLifetime: 0,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", addr)
		},
	}
}
