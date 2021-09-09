package service

import (
	"bala/app/config"
	"bala/app/db"
	"bala/app/db/mysql"
	"bala/app/model/entity"
	"bala/app/model/repo"
	"bala/app/myerr"
	"bala/app/protocol"
	"bala/app/protocol/agent"
	"bala/app/protocol/errcode"
	"strconv"

	"github.com/google/wire"
)

var AccountSet = wire.NewSet(wire.Struct(new(Account), "*"))

type Account struct {
	Config    *config.Config
	DBManager *db.Manager
	RoleRepo  *repo.RoleRepository
	AccRepo   *repo.AccountRepository
}

// Create 创建账号
func (a *Account) Create(req *protocol.AccountCreateRequest) error {
	commonDB := a.DBManager.GetCommonDB()
	err := commonDB.Transaction(func(ctx *mysql.DB) error {
		acc := &entity.Account{
			Account:      req.DevId,
			IMEI:         req.IMEI,
			RegisteredIP: req.AccessIP,
		}
		// 查找账号
		if exist, err := a.AccRepo.Find(ctx, acc); exist {
			return myerr.NewFmt(errcode.AccountCreateFail, "create account : %s already exist", acc.Account)
		} else if err != nil {
			return myerr.New(errcode.AccountCreateFail, err)
		}
		// 创建账号
		err := a.AccRepo.Create(ctx, acc)
		if err != nil {
			return myerr.New(errcode.AccountCreateFail, err)
		}
		// 获取Id
		roleId := a.getRoleId(acc.Id)
		dbId := a.getDbId(acc.Id)
		// 创建角色
		var role *entity.Role
		gameDB := a.DBManager.GetGameDB(dbId)
		err = gameDB.Transaction(func(gtx *mysql.DB) error {
			role = &entity.Role{
				RoleId:  roleId,
				DbId:    dbId,
				Account: acc.Account,
			}
			return a.RoleRepo.Create(gtx, role)
		})
		return err
	})
	return err
}

// Login 验证成功后返回常用数据
func (a *Account) Login(req *protocol.AccountAuthRequest) (*protocol.AccountAuthResponse, error) {
	resp := &protocol.AccountAuthResponse{}
	resp.CurrentVersion = 1
	resp.MinimumVersion = 1
	resp.IsDevelopment = true
	resp.UpdateRequired = true
	// 账号
	commonDB := a.DBManager.GetCommonDB()
	err := commonDB.Transaction(func(tx *mysql.DB) error {
		acc := &entity.Account{Account: req.DevId}
		if exist, err := a.AccRepo.Find(tx, acc); !exist {
			return myerr.NewFmt(errcode.AccountAuthNotCreated, "unknown account: %s", acc.Account)
		} else if err != nil {
			return myerr.New(errcode.AccountAuthNotCreated, err)
		}
		// 获取Id
		roleId := a.getRoleId(acc.Id)
		dbId := a.getDbId(acc.Id)
		// 角色
		gameDB := a.DBManager.GetGameDB(dbId)
		return gameDB.Transaction(func(tx *mysql.DB) error {
			role, err := a.RoleRepo.Find(tx, roleId)
			if err != nil {
				return myerr.New(errcode.None, err)
			}
			if role == nil {
				return myerr.NewFmt(errcode.None, "login find role is null account is: %s", role.Account)
			}
			// 基本信息
			resp.AccountDB = &agent.AccountDB{
				ServerId:        role.DbId,
				RoleId:          role.RoleId,
				Nickname:        role.NickName,
				DevId:           acc.Account,
				State:           acc.Type,
				Level:           role.Level,
				Exp:             role.Exp,
				VIPLevel:        role.VipLevel,
				CreateDate:      acc.CreateAt,
				LastConnectTime: role.LastUpdateAt,
			}
			// 考勤奖励
			// resp.AttendanceBookRewards =
			// 历史考勤
			// resp.AttendanceHistoryDBs =
			// 功能开放
			// resp.OpenConditions =
			// 背包道具
			// resp.MonthlyProductParcel =
			// 邮件
			// resp.MonthlyProductMail =
			return a.RoleRepo.Update(tx, role, "last_login")
		})
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// getDbId 生成db的id
func (a *Account) getDbId(n int) int {
	if n < 1 {
		panic("generateDbId error:" + strconv.Itoa(n))
	}
	database := n % a.DBManager.GetGameDBCount()
	if database == 0 {
		database = a.DBManager.GetGameDBCount()
	}
	table := n % a.Config.Get().Mysql.Tables
	if table == 0 {
		table = a.Config.Get().Mysql.Tables
	}
	return 1000000 + database*1000 + table
}

// getRoleId 生成role的id
func (a *Account) getRoleId(n int) int {
	return 1000000000 + n
}
