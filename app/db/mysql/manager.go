package mysql

import (
	"bala/app/config"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	commons = make([]interface{}, 0)
	games   = make([]interface{}, 0)
	tables  int
)

type Manager struct {
	common *DB
	dbs    []*DB // 可以增加新的连接,不允许改变已有的连接
	dbsMux *sync.RWMutex
}

func (m *Manager) Count() int {
	m.dbsMux.RLock()
	defer m.dbsMux.RUnlock()
	return len(m.dbs)
}

func (m *Manager) GetCommonDB() *DB {
	return m.common
}

func (m *Manager) GetGameDB(index int) (*DB, error) {
	m.dbsMux.RLock()
	defer m.dbsMux.RUnlock()
	if index < 0 || index >= len(m.dbs) {
		return nil, fmt.Errorf("mysql GetGameDB index is out of bounds: %d", index)
	}
	return m.dbs[index], nil
}

func RegisterCommonEntity(v interface{}) {
	commons = append(commons, v)
}

func RegisterGameEntity(v interface{}) {
	games = append(games, v)
}

func InitMysqlManager(c *config.MysqlShards) (*Manager, error) {
	m := &Manager{dbsMux: new(sync.RWMutex)}
	tables = c.Tables
	// common
	db, err := initMysql(c.Debug, c.Common)
	if err != nil {
		return nil, err
	}
	if c.DDL {
		err = createTable(db, 1, commons...)
		if err != nil {
			return nil, err
		}
	}
	m.common = db

	m.dbsMux.Lock()
	// game
	for _, v := range c.Shards {
		db, err := initMysql(c.Debug, v)
		if err != nil {
			return nil, err
		}
		if c.DDL {
			err = createTable(db, tables, games...)
			if err != nil {
				return nil, err
			}
		}
		m.dbs = append(m.dbs, db)
	}
	m.dbsMux.Unlock()
	return m, nil
}

func initMysql(debug bool, c *config.MysqlConfig) (*DB, error) {
	db, err := gorm.Open(mysql.Open(c.Dsn), &gorm.Config{
		PrepareStmt:            false, // 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
		SkipDefaultTransaction: true,  // 默认情况下，GORM在事务中执行单一的创建、更新、删除操作，以确保数据库数据的完整性，设置' SkipDefaultTransaction '为true来禁用它
		DisableAutomaticPing:   false, // 在完成初始化后，GORM 会自动 ping 数据库以检查数据库的可用性，若要禁用该特性，可将其设置为 true
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(c.MaxIdleCount)
	// 最大连接数
	sqlDB.SetMaxOpenConns(c.MaxOpen)
	// 连接最大存活时间
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxIdleTime) * time.Minute)
	// 空闲连接最大存活时间
	sqlDB.SetConnMaxIdleTime(time.Duration(c.MaxIdleTime) * time.Minute)
	//
	if debug {
		db = db.Debug()
	}
	return &DB{DB: db}, nil
}

func createTable(db *DB, n int, values ...interface{}) error {
	migrator := db.Migrator()
	for _, v := range values {
		typeof := reflect.TypeOf(v)
		name := strings.ToLower(typeof.Elem().Name())
		field := reflect.ValueOf(migrator).
			FieldByName("DB").
			Elem().FieldByName("Statement").
			Elem().FieldByName("Table")
		// 自动建表
		for i := 1; i <= n; i++ {
			var tableName string
			if n == 1 {
				tableName = name
			} else {
				tableName = fmt.Sprintf("%s_%d", name, i)
			}
			field.Set(reflect.ValueOf(tableName))
			if !migrator.HasTable(v) {
				err := migrator.CreateTable(v)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
