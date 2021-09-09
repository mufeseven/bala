package repo

import "github.com/google/wire"

var WireSet = wire.NewSet(
	AccountSet,
	RoleSet,
	CurrencySet,
	ItemSet,
	AttendanceSet,
	CharacterSet,
)
