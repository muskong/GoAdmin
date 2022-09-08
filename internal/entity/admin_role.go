package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	AdminRole struct {
		gorm.Model
		ParentUuid  string          `json:"ParentUuid,omitempty" db:"parent_uuid"`
		Name        string          `json:"Name,omitempty" db:"name"`
		Rules       gorm.JsonString `json:"Rules,omitempty" db:"rules"`
		Description string          `json:"Description,omitempty" db:"description"`
		State       string          `json:"State,omitempty" db:"state"`
	}
	role struct{}
)

var Role = &role{}

func (e *role) db() *gdb.DB {
	return gorm.NewModel(&AdminRole{}).Where("deleted_at IS NULL")
}

func (e *role) GetRole(roleId string) (*AdminRole, error) {
	var role AdminRole
	err := e.db().Where("uuid = ?", roleId).First(&role).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &role, err
}
func (e *role) GetRoleAll() (roles []*AdminRole, err error) {
	err = e.db().Find(&roles).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (e *role) GetRolesByNames(roleName []string) (roles []*AdminRole, err error) {
	err = e.db().Where("name IN ?", roleName).Order("id desc").Find(&roles).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

// 按选项查询集合
// offset 跳过
// limit 读取数量
// field  排序字段
// sort  排序   DESC 倒叙 ， ASC 正序
func (e *role) GetRoles(offset, limit int) (roles []*AdminRole, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&roles).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (e *role) InsertAdminRole(role *AdminRole) (err error) {
	err = e.db().Create(role).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (e *role) UpdateAdminRole(role *AdminRole) (err error) {
	err = e.db().Updates(role).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (e *role) DeleteAdminRole(roleId string) error {
	ar := AdminRole{}
	ar.DeletedAt = gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := e.db().Where("uuid = ?", roleId).Updates(ar).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *role) UpdateAdminRoleRules(roleId string, ruleIds []string) error {
	err := e.db().Where("uuid = ?", roleId).Update("rules", ruleIds).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}
