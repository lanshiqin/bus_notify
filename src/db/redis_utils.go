package db

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

func SetString(key interface{}, value interface{}) error {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return errors.New("connect to redis error")
	}
	defer c.Close()
	_, err = c.Do("SET", key, value)
	if err != nil {
		return errors.New("redis set failed")
	}
	return nil
}

func GetString(key interface{}) string {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		logs.Error("connect to redis error")
		return ""
	}
	defer c.Close()
	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		logs.Error("redis get failed")
		return ""
	}
	return value
}
