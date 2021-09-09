package entity

import (
	"bala/app/db/mysql"
	"time"
)

type Account struct {
	Id           int       `gorm:"column:id;primaryKey;auto_increment"`
	Account      string    `gorm:"column:account;type:varchar(30);index:idx_account;unique"`
	Type         int       `gorm:"column:type;type:int(11)"`
	IMEI         int64     `gorm:"column:imei;type:bigint(20)"`
	CreateAt     time.Time `gorm:"column:create_at;autoCreateTime"`
	RegisteredIP string    `gorm:"column:registered_ip;type:varchar(20)"`
}

func init() {
	mysql.RegisterCommonEntity(&Account{})
}
