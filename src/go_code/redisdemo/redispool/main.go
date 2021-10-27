package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个全局pool
var pool *redis.Pool

//当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		TestOnBorrow:    nil,
		MaxIdle:         0, //最大空闲连接数
		MaxActive:       0, //数据库最大连接数，0为无限制
		IdleTimeout:     0, //最大空闲时间
		Wait:            false,
		MaxConnLifetime: 0,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	//从pool中取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "monster", "wolf~~")
	if err != nil {
		fmt.Println("conn.do() err: ", err)
	}

	r, err := redis.String(conn.Do("get", "monster"))
	if err != nil {
		fmt.Println("conn.do() err: ", err)
	}
	fmt.Println("r=", r)

	//pool.Close()

}
