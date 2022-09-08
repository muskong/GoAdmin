package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/idworker"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		WechatOpenid  string          `json:"WechatOpenid" db:"wechat_openid"`
		Name          string          `json:"Name" db:"name"`
		Avatar        string          `json:"Avatar" db:"avatar"`
		NickName      string          `json:"NickName" db:"nick_name"`
		Password      string          `json:"Password" db:"password"`
		PayPassword   string          `json:"PayPassword" db:"pay_password"`
		AccountAmount string          `json:"AccountAmount" db:"account_amount"`
		Groups        gorm.JsonString `json:"Groups" db:"groups"`

		Verified UserVerified `gorm:"foreignkey:UserUuid;references:Uuid"`
		Bank     UserBank     `gorm:"foreignkey:UserUuid;references:Uuid"`
	}
	_user struct{}
)

var UserEntity = new(_user)

func (*_user) db() *gdb.DB {
	return gorm.NewModel(&User{}).Preload("Verified").Preload("Bank")
}

func (e *_user) GetUser(uuid string) (*User, error) {
	var data User
	err := e.db().Where("uuid = ?", uuid).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_user) GetUsers(offset, limit int) (users []*User, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_user) Insert(user *User) (err error) {
	user.Uuid = idworker.StringNanoid(30)

	err = e.db().Create(user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_user) Delete(userId int) error {
	err := e.db().Delete(&User{}, userId).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_user) GetAll() (users []*User, err error) {
	err = e.db().Order("id desc").Find(&users).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
