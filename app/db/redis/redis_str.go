package redis

import (
	"bala/app/db/redis/serializable"
	"bala/app/log"
	"time"

	"github.com/gomodule/redigo/redis"
)

func (d *DB) GetString(key interface{}) (s string, err error) {
	conn := d.Pool.Get()
	defer d.close(conn)
	s, err = redis.String(conn.Do("GET", key))
	if err != nil {
		log.Local().Errorf("redis command GetString err: %s", err)
	}
	return
}

func (d *DB) SetString(key, value interface{}, ttl ...time.Duration) (reply interface{}, err error) {
	conn := d.Pool.Get()
	defer d.close(conn)
	if len(ttl) < 1 {
		reply, err = conn.Do("SET", key, value)
	} else {
		reply, err = conn.Do("SET", key, value, "EX", ttl[0].Seconds())
	}
	if err != nil {
		log.Local().Errorf("redis command SetString err: %s", err)
	}
	return
}

func (d *DB) Get(key, ref interface{}) (err error) {
	conn := d.Pool.Get()
	defer d.close(conn)
	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		log.Local().Errorf("redis command Get err: %s", err)
		return
	}
	ins := serializable.GetInstance()
	err = ins.Deserialization(reply, ref)
	if err != nil {
		log.Local().Errorf("redis command Get err: %s", err)
	}
	return
}

func (d *DB) Set(key, value interface{}, ttl ...time.Duration) (reply interface{}, err error) {
	conn := d.Pool.Get()
	defer d.close(conn)
	ins := serializable.GetInstance()
	bytes, err := ins.Serialization(value)
	if err != nil {
		log.Local().Errorf("redis command Set err: %s", err)
		return nil, err
	}
	if len(ttl) < 1 {
		reply, err = conn.Do("SET", key, bytes)
	} else {
		reply, err = conn.Do("SET", key, bytes, "EX", ttl[0].Seconds())
	}
	if err != nil {
		log.Local().Errorf("redis command Set err: %s", err)
	}
	return
}
