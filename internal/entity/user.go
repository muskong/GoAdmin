package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
)

type (
	AdminUser struct {
		/** ID **/
		Id int `json:"Id,omitempty" db:"id"`
		/** 用户名 **/
		Name string `json:"Name,omitempty" db:"name"`
		/** 密码 **/
		Password string `json:"Password,omitempty" db:"password"`
		/** 角色 **/
		Roles gorm.JsonString `json:"Roles,omitempty" db:"roles"`

		/** 发布时间 **/
		CreatedAt gorm.TimeString `json:"CreatedAt,omitempty" db:"created_at"`
		/** 更新时间 **/
		UpdatedAt gorm.TimeString `json:"UpdatedAt,omitempty" db:"updated_at"`
		DeletedAt gorm.TimeString `json:"DeletedAt,omitempty" db:"deleted_at"`
	}
	user struct{}
)

var User = &user{}

func (*user) GetAdminName(name string) (*AdminUser, error) {
	db := gorm.ClientNew().Model(AdminUser{}).Where("deleted_at IS NULL")

	var user AdminUser
	err := db.Where("name = ?", name).First(&user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &user, err
}

func (*user) GetUser(adminId int) (*AdminUser, error) {
	db := gorm.ClientNew().Model(AdminUser{}).Where("deleted_at IS NULL")
	var user AdminUser
	err := db.Where("id = ?", adminId).First(&user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &user, err
}

// 按选项查询集合
func (*user) GetUsers(page, limit int) (users []*AdminUser, count int64, err error) {
	db := gorm.ClientNew().Model(AdminUser{}).Where("deleted_at IS NULL")
	err = db.Count(&count).Order("id desc").Limit(limit).Offset(page).Find(&users).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (*user) InsertAdminUser(user *AdminUser) (*AdminUser, error) {
	db := gorm.ClientNew().Model(AdminUser{}).Where("deleted_at IS NULL")
	err := db.Create(user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return user, err
}

func (*user) UpdateAdminUser(user *AdminUser) (*AdminUser, error) {
	db := gorm.ClientNew().Model(AdminUser{}).Where("deleted_at IS NULL")
	err := db.Updates(user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return user, err
}

func (*user) DeleteAdminUser(userId int) error {
	db := gorm.ClientNew().Model(&AdminUser{}).Where("deleted_at IS NULL")
	deletedAt := gorm.TimeString(time.Now().Format("2006-01-02 15:04:05"))
	err := db.Where("id=?", userId).Updates(AdminUser{DeletedAt: deletedAt}).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}
