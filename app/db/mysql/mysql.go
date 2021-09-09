package mysql

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
	tn *int // table number
}

func NewDB(db *gorm.DB, tn *int) *DB {
	return &DB{DB: db.Session(&gorm.Session{NewDB: true}), tn: tn}
}

// Transaction 开启事务
func (db *DB) Transaction(fc func(tx *DB) error, opts ...*sql.TxOptions) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		return fc(NewDB(tx, db.tn))
	}, opts...)
}

// Table 指定entity映射的表,主要为了gameDB的使用
func (db *DB) Table(name string) (tx *gorm.DB) {
	if db.tn == nil {
		return db.DB.Table(name)
	} else {
		return db.DB.Table(fmt.Sprintf("%s_%d", name, *db.tn))
	}
}

func (db *DB) TableObj(v interface{}) (tx *gorm.DB) {
	if db.tn == nil {
		return db.DB.Table(getTableName(v))
	} else {
		return db.DB.Table(fmt.Sprintf("%s_%d", getTableName(v), *db.tn))
	}
}

func getTableName(v interface{}) string {
	name := reflect.TypeOf(v).Elem().Name()
	return strings.ToLower(name)
}
