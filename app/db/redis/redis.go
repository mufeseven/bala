package redis

import (
	"bala/app/log"
	"time"

	"github.com/gomodule/redigo/redis"

	_ "bala/app/db/redis/serializable/json"
)

type DB struct {
	Pool *redis.Pool
}

func (d *DB) close(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		log.Local().Errorf("redis conn close err: %s", err)
	}
}

// Select 切换到指定的数据库，数据库索引号 index 用数字值指定，以 0 作为起始索引值
func (d *DB) Select(index int) (b bool, err error) {
	conn := d.Pool.Get()
	defer d.close(conn)
	b, err = redis.Bool(conn.Do("SELECT ", index))
	if err != nil {
		log.Local().Errorf("redis command Select err: %s", err)
	}
	return
}

// Exists 检查给定 key 是否存在
func (d *DB) Exists(key interface{}) (b bool, err error) {
	conn := d.Pool.Get()
	defer d.close(conn)
	b, err = redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		log.Local().Errorf("redis command Exists err: %s", err)
	}
	return
}

// Expire 设置生存时间
func (d *DB) Expire(key interface{}, ttl time.Duration) (i int, err error) {
	conn := d.Pool.Get()
	defer d.close(conn)
	i, err = redis.Int(conn.Do("EXPIRE", key, ttl.Seconds()))
	if err != nil {
		log.Local().Errorf("redis command Expire err: %s", err)
	}
	return
}

// Ttl 以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)
func (d *DB) Ttl(key interface{}) (i int, err error) {
	conn := d.Pool.Get()
	defer d.close(conn)
	i, err = redis.Int(conn.Do("TTL", key))
	if err != nil {
		log.Local().Errorf("redis command Ttl err: %s", err)
	}
	return
}

// Persist 移除给定 key 的生存时间，将这个 key 从『易失的』(带生存时间 key )转换成『持久的』(一个不带生存时间、永不过期的 key )
func (d *DB) Persist(key interface{}) (i int, err error) {
	conn := d.Pool.Get()
	defer d.close(conn)
	i, err = redis.Int(conn.Do("PERSIST", key))
	if err != nil {
		log.Local().Errorf("redis command Persist err: %s", err)
	}
	return
}

// Del 删除给定的一个或多个 key
func (d *DB) Del(key ...interface{}) (i int, err error) {
	conn := d.Pool.Get()
	defer d.close(conn)
	i, err = redis.Int(conn.Do("DEL", key...))
	if err != nil {
		log.Local().Errorf("redis command Del err: %s", err)
	}
	return
}
