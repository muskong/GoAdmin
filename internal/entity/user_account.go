package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/idworker"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	UserAccount struct {
		// gdb.Model
		ID       int     `json:"ID,omitempty" db:"id"`
		UserUuid string  `json:"UserUuid,omitempty" db:"user_uuid"`
		Before   float64 `json:"Before,omitempty" db:"before"`
		Change   float64 `json:"Change,omitempty" db:"change"`
		After    float64 `json:"After,omitempty" db:"after"`
		Remark   string  `json:"Remark,omitempty" db:"remark"`
		Table    string  `json:"Table,omitempty" db:"table"`
		TableId  string  `json:"TableId,omitempty" db:"table_id"`

		CreatedAt gorm.TimeString `json:"CreatedAt,omitempty" db:"created_at"`
	}
	_userAccount struct{}
)

var UserAccountEntity = new(_userAccount)

func (*_userAccount) db() *gdb.DB {
	return gorm.NewModel(&UserAccount{}).Omit("UpdatedAt", "DeletedAt")
}

func (e *_userAccount) GetUserAccount(uuid string) (*UserAccount, error) {
	var data UserAccount
	err := e.db().Where("uuid = ?", uuid).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_userAccount) GetUserAccounts(offset, limit int) (userAccounts []*UserAccount, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&userAccounts).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_userAccount) Insert(userAccount *UserAccount) (err error) {
	userAccount.UserUuid = idworker.StringNanoid(30)

	err = e.db().Create(userAccount).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_userAccount) GetAll() (userAccounts []*UserAccount, err error) {
	err = e.db().Order("id desc").Find(&userAccounts).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
