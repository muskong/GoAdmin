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
		ID       int     `json:"ID" db:"id"`
		UserUuid string  `json:"UserUuid" db:"user_uuid"`
		Before   float64 `json:"Before" db:"before"`
		Change   float64 `json:"Change" db:"change"`
		After    float64 `json:"After" db:"after"`
		Remark   string  `json:"Remark" db:"remark"`
		Table    string  `json:"Table" db:"table"`
		TableId  string  `json:"TableId" db:"table_id"`

		CreatedAt gorm.TimeString `json:"CreatedAt" db:"created_at"`

		Verified UserVerified `gorm:"foreignkey:UserUuid;references:UserUuid"`
		User     User         `gorm:"foreignkey:Uuid;references:UserUuid"`
	}
	_userAccount struct{}
)

var UserAccountEntity = new(_userAccount)

func (*_userAccount) db() *gdb.DB {
	return gorm.NewModel(&UserAccount{}).Omit("UpdatedAt", "DeletedAt").Preload("User").Preload("Verified")
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
