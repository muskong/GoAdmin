package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/idworker"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	UserGroup struct {
		gdb.Model
		ID        int             `json:"ID,omitempty" db:"id"`
		GroupUuid string          `json:"GroupUuid,omitempty" db:"group_uuid"`
		Title     string          `json:"Title,omitempty" db:"title"`
		Content   string          `json:"Content,omitempty" db:"content"`
		CreatedAt gorm.TimeString `json:"CreatedAt,omitempty" db:"created_at"`
		UpdatedAt gorm.TimeString `json:"UpdatedAt,omitempty" db:"updated_at"`
		DeletedAt gorm.NullString `json:"DeletedAt,omitempty" db:"deleted_at"`
	}
	_userGroup struct{}
)

var UserGroupEntity = new(_userGroup)

func (*_userGroup) db() *gdb.DB {
	return gorm.NewModel(&UserGroup{})
}

func (e *_userGroup) GetUserGroup(uuid string) (*UserGroup, error) {
	var data UserGroup
	err := e.db().Where("uuid = ?", uuid).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_userGroup) GetUserGroups(offset, limit int) (userGroups []*UserGroup, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&userGroups).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_userGroup) Insert(userGroup *UserGroup) (err error) {
	userGroup.GroupUuid = idworker.StringNanoid(30)

	err = e.db().Create(userGroup).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (e *_userGroup) Delete(groupId int) error {
	err := e.db().Delete(&UserGroup{}, groupId).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_userGroup) GetAll() (userGroups []*UserGroup, err error) {
	err = e.db().Order("id desc").Find(&userGroups).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
