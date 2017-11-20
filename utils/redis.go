package utils

import (
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	client redis.Conn
}

//New redis conn
func NewRedis(host string, port string, db int, password string) *Redis {
	c, err := redis.Dial("tcp", host+":"+port, redis.DialDatabase(db), redis.DialPassword(password))
	if err != nil {
		panic(err)
	}
	r := &Redis{client: c}
	return r
}

//Set string
func (c *Redis) Set(key string, value string, expire int64) (err error) {
	_, err = c.client.Do("SET", key, value)
	if expire > 0 {
		c.client.Do("EXPIRE", key, expire)
	}
	return err
}

//Get string
func (c *Redis) Get(key string) (value string, err error) {
	value, err = redis.String(c.client.Do("GET", key))
	return value, err
}

//Set hash
func (c *Redis) HSet(key string, value interface{}, expire int64) (err error) {
	_, err = c.client.Do("HMSET", redis.Args{}.Add(key).AddFlat(value)...)
	if expire > 0 {
		c.client.Do("EXPIRE", key, expire)
	}
	return err
}

//Get hash
func (c *Redis) HGet(key string, value interface{}) (err error) {
	v, err := redis.Values(c.client.Do("HGETALL", key))
	if err != nil {
		return err
	}
	if err := redis.ScanStruct(v, value); err != nil {
		return err
	}
	return err
}

//Del redis data
func (c *Redis) Del(key string) (err error) {
	_, err = c.client.Do("DEL", key)
	return err
}

//Is exist or not
func (c *Redis) IsExist(key string) (flag bool, err error) {
	flag, err = redis.Bool(c.client.Do("EXISTS", key))
	return flag, err
}

//Close redis conn
func (c *Redis) CloseRedis() {
	c.client.Close()
}
