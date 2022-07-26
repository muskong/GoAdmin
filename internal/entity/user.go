package entity

import (
	"database/sql"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
)

type (
	AdminUser struct {
		/** ID **/
		Id int64 `json:"Id,omitempty" db:"id"`
		/** 用户名 **/
		Name string `json:"Name,omitempty" db:"name"`
		/** 密码 **/
		Password string `json:"Password,omitempty" db:"password"`
		/** 角色 **/
		Roles gorm.JsonString `json:"Roles,omitempty" db:"roles"`

		/** 发布时间 **/
		CreatedAt string `json:"createdAt,omitempty" db:"created_at"`
		/** 更新时间 **/
		UpdatedAt sql.NullString `json:"updatedAt,omitempty" db:"updated_at"`
		DeletedAt sql.NullString `json:"deletedAt,omitempty" db:"deleted_at"`
	}
)

func GetAdminName(name string) (*AdminUser, error) {
	db := gorm.ClientNew().Model(AdminUser{}).Where("deleted_at IS NULL")

	var user AdminUser
	err := db.Where("name = ?", name).First(&user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &user, err
}

func GetUser(adminId int64) (*AdminUser, error) {
	db := gorm.ClientNew().Model(AdminUser{}).Where("deleted_at IS NULL")
	var user AdminUser
	err := db.Where("id = ?", adminId).First(&user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &user, err
}

// 按选项查询集合
func GetUsers(page, limit int) (users []*AdminUser, count int64, err error) {
	db := gorm.ClientNew().Model(AdminUser{}).Where("deleted_at IS NULL")
	err = db.Count(&count).Order("id desc").Limit(limit).Offset(page).Find(&users).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func InsertAdminUser(user *AdminUser) (*AdminUser, error) {
	db := gorm.ClientNew().Model(AdminUser{}).Where("deleted_at IS NULL")
	err := db.Create(user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return user, err
}
