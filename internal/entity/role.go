package entity

import (
	"database/sql"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
)

type (
	AdminRole struct {
		Id          int64          `json:"Id,omitempty" db:"id"`
		Name        string         `json:"Name,omitempty" db:"name"`
		Rules       gorm.JsonInt64 `json:"Rules,omitempty" db:"rules"`
		Description string         `json:"Description,omitempty" db:"description"`
		CreatedAt   string         `json:"createdAt,omitempty" db:"created_at"`
		UpdatedAt   sql.NullString `json:"updatedAt,omitempty" db:"updated_at"`
		DeletedAt   sql.NullString `json:"deletedAt,omitempty" db:"deleted_at"`
	}
)

func GetRole(roleId int64) (*AdminRole, error) {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	var role AdminRole
	err := db.Where("id = ?", roleId).First(&role).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &role, err
}
func GetRoleAll() (roles []*AdminRole, err error) {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	err = db.Find(&roles).Error
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
func GetRoles(offset, limit int) (roles []*AdminRole, count int64, err error) {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	err = db.Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&roles).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func InsertAdminRole(role *AdminRole) (err error) {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	err = db.Create(role).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func UpdateAdminRoleRules(roleId int64, ruleIds []int64) error {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	err := db.Where("id=?", roleId).Update("rules", ruleIds).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}
