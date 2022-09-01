package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	AdminUser struct {
		/** ID **/
		Id     int    `json:"Id,omitempty" db:"id"`
		Nanoid string `json:"Nanoid,omitempty" db:"nanoid"`
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
		DeletedAt gorm.NullString `json:"DeletedAt,omitempty" db:"deleted_at"`
	}
	_adminUser struct{}
)

var AdminUserEntity = &_adminUser{}

func (e *_adminUser) db() *gdb.DB {
	return gorm.NewModel(&AdminUser{}).Where("deleted_at IS NULL")
}

func (e *_adminUser) GetAdminName(name string) (*AdminUser, error) {

	var user AdminUser
	err := e.db().Where("name = ?", name).First(&user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &user, err
}

func (e *_adminUser) GetUser(adminId int) (*AdminUser, error) {
	var user AdminUser
	err := e.db().Where("id = ?", adminId).First(&user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &user, err
}

// 按选项查询集合
func (e *_adminUser) GetUsers(page, limit int) (users []*AdminUser, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(page).Find(&users).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_adminUser) InsertAdminUser(user *AdminUser) (*AdminUser, error) {
	err := e.db().Create(user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return user, err
}

func (e *_adminUser) UpdateAdminUser(user *AdminUser) (*AdminUser, error) {
	err := e.db().Updates(user).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return user, err
}

func (e *_adminUser) DeleteAdminUser(userId int) error {
	au := AdminUser{}
	au.DeletedAt = gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := e.db().Where("id=?", userId).Updates(au).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}
