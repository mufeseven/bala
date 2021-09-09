package redis

import (
	"bala/app/config"
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Manager struct {
	dbs    []*DB // 可以增加新的连接,不允许改变已有的连接
	dbsMux *sync.RWMutex
}

func (m *Manager) Count() int {
	m.dbsMux.RLock()
	defer m.dbsMux.RUnlock()
	return len(m.dbs)
}

func (m *Manager) GetCacheDB(index int) (*DB, error) {
	m.dbsMux.RLock()
	defer m.dbsMux.RUnlock()
	if index < 0 || index >= len(m.dbs) {
		return nil, fmt.Errorf("redis GetCacheDB index is out of bounds: %d", index)
	}
	return m.dbs[index], nil
}

func InitRedisManager(c *config.RedisShards) (*Manager, error) {
	m := &Manager{dbsMux: new(sync.RWMutex)}
	m.dbsMux.Lock()
	for _, v := range c.Shards {
		db, err := initRedis(v)
		if err != nil {
			return nil, err
		}
		m.dbs = append(m.dbs, db)
	}
	m.dbsMux.Unlock()
	return m, nil
}

func initRedis(c *config.RedisConfig) (*DB, error) {
	options := []redis.DialOption{
		redis.DialConnectTimeout(time.Duration(c.ConnectTimeout) * time.Second),
		redis.DialKeepAlive(time.Duration(c.KeepAlive) * time.Minute),
		redis.DialReadTimeout(time.Duration(c.ReadTimeout) * time.Millisecond),
		redis.DialWriteTimeout(time.Duration(c.WriteTimeout) * time.Millisecond),
	}
	// 连接测试
	conn, err := redis.DialURL(c.Url, options...)
	if err != nil {
		return nil, err
	}
	err = conn.Close()
	if err != nil {
		return nil, err
	}
	// 配置连接池
	p := &redis.Pool{
		MaxIdle:         c.MaxIdle,
		MaxActive:       c.MaxActive,
		IdleTimeout:     time.Duration(c.IdleTimeout) * time.Minute,
		MaxConnLifetime: time.Duration(c.MaxLifetime) * time.Minute,
		// If Wait is true and the pool is at the MaxActive limit, then Get() waits
		// for a connection to be returned to the pool before returning.
		Wait: true,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(c.Url, options...)
		},
	}
	return &DB{Pool: p}, nil
}
