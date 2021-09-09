package entity

import "time"

// 额外附加
type Extra struct {
	Version      int16     `gorm:"column:version;default:0"`
	CreateAt     time.Time `gorm:"column:create_at;autoCreateTime"`
	LastUpdateAt time.Time `gorm:"column:last_update_at;autoUpdateTime"`
}
