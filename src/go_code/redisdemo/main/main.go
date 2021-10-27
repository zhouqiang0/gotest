package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go 向redis写入数据和读取数据
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()
	//写入数据 string[key-val]
	_, err = conn.Do("set", "name", "tom猫")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//r是空接口interface{},对应值是string，需要转换redis.String()
	r1, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Println("op suc", r1)

	//写入数据 string[key-val]
	_, err = conn.Do("hmset", "hash1", "hname", "jerry老鼠", "age", 12)
	if err != nil {
		fmt.Println("hash set err=", err)
		return
	}

	//r是空接口interface{},对应值是string，需要转换redis.String()
	r2, err := redis.Strings(conn.Do("hgetall", "hash1"))
	if err != nil {
		fmt.Println("hash get err=", err)
		return
	}
	fmt.Println("op hash suc", r2)
}
