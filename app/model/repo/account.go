package repo

import (
	"bala/app/db/mysql"
	"bala/app/model/entity"

	"github.com/google/wire"
)

var AccountSet = wire.NewSet(wire.Struct(new(AccountRepository), "*"))

type AccountRepository struct{}

// Find 查找账号
func (a *AccountRepository) Find(db *mysql.DB, acc *entity.Account) (bool, error) {
	// db.Where(&entity.Account{Account: acc.Account}).Find(acc)
	// db.Where(map[string]interface{}{"account": acc.Account}).Find(&acc)
	// db.Where(acc, "account").Find(acc)
	result := db.Where("account = ?", acc.Account).Find(acc)
	return result.RowsAffected > 0, result.Error
}

// Create 创建账号
func (a *AccountRepository) Create(db *mysql.DB, acc *entity.Account) error {
	return db.Create(acc).Error
}
