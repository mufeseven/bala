package db

import (
	"bala/app/config"
	"bala/app/db/mysql"
	"bala/app/db/redis"
)

type Manager struct {
	tables int
	mysql  *mysql.Manager
	redis  *redis.Manager
}

func InitManager(c *config.Config) (*Manager, func(), error) {
	s := &Manager{}
	var err error
	s.tables = c.Get().Mysql.Tables
	s.mysql, err = mysql.InitMysqlManager(c.Get().Mysql)
	if err != nil {
		return nil, nil, err
	}
	s.redis, err = redis.InitRedisManager(c.Get().Redis)
	if err != nil {
		return nil, nil, err
	}
	return s, func() {}, nil
}

func (m *Manager) GetCommonDB() *mysql.DB {
	db := m.mysql.GetCommonDB()
	return mysql.NewDB(db.DB, nil)
}

func (m *Manager) GetGameDB(dbId int) *mysql.DB {
	dbId = dbId - 1000000
	database := dbId / 1000
	db, err := m.mysql.GetGameDB(database - 1)
	if err != nil {
		return nil
	}
	if m.tables <= 1 {
		return mysql.NewDB(db.DB, nil)
	}
	tn := dbId - database*1000
	return mysql.NewDB(db.DB, &tn)
}

func (m *Manager) GetCacheDB(dbId int) *redis.DB {
	i := dbId % m.redis.Count()
	db, err := m.redis.GetCacheDB(i)
	if err != nil {
		return nil
	}
	return db
}

func (m *Manager) GetGameDBCount() int {
	return m.mysql.Count()
}

func (m *Manager) GetCacheDBCount() int {
	return m.redis.Count()
}
