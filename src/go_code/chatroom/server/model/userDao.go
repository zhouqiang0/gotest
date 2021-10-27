package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"test/src/go_code/chatroom/common/message"
)

//userDao完成对User结构体的各种操作
//在服务器启动时(main.go)，初始化一个userDao,将其作为全局变量
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{pool: pool}
	return userDao
}

//1.根据用户id，返回一个User实例
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过id去redis查询用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil { //表示users中无该id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//将res反序列化为User实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
}

//完成登录校验
//Login :完成对用户的验证，正确返回user实例，否则返回err
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从userDao中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//err == nil 证明用户是能获取的
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

//Register :完成对用户的验证，正确返回user实例，否则返回err
func (this *UserDao) Register(user *message.User) (err error) {
	//先从userDao中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()

	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS //能查到说明该用户已经存在
		return
	}
	//接下来完成注册
	dbData, err := json.Marshal(user)
	if err != nil {
		return
	}
	//入库
	_, err = conn.Do("hset", "users", user.UserId, string(dbData))
	if err != nil {
		fmt.Println("保存注册用户错误 err : ", err)
		return
	}
	return
}
