package agent

import (
	"time"
)

type AccountDB struct {
	ServerId                   int
	RoleId                     int // 为重构增加字段
	Nickname                   string
	CallName                   string
	DevId                      string
	State                      int // excel.AccountStateNormal
	Level                      int
	Exp                        int64
	Comment                    string
	LobbyMode                  int
	RepresentCharacterServerId int
	MemoryLobbyUniqueId        int64
	LastConnectTime            time.Time
	BirthDay                   time.Time
	CallNameUpdateTime         time.Time
	PublisherAccountId         int64
	RetentionDays              int
	VIPLevel                   int
	CreateDate                 time.Time
	UnReadMailCount            int
	LinkRewardDateDate         time.Time
}

type AccountCurrencyDB struct {
	CurrencyDict   map[int]int64     // key:excel.CurrencyTypesInvalid, value:count
	UpdateTimeDict map[int]time.Time // key:excel.CurrencyTypesInvalid, value:time
}
