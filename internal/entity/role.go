package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/idworker"
	"github.com/muskong/GoPkg/zaplog"
)

type (
	AdminRole struct {
		Id           int             `json:"Id,omitempty" db:"id"`
		Nanoid       string          `json:"Nanoid,omitempty" db:"nanoid"`
		ParentNanoid string          `json:"ParentNanoId,omitempty" db:"parent_nanoid"`
		Name         string          `json:"Name,omitempty" db:"name"`
		Rules        gorm.JsonString `json:"Rules,omitempty" db:"rules"`
		Description  string          `json:"Description,omitempty" db:"description"`
		State        string          `json:"State,omitempty" db:"state"`
		CreatedAt    gorm.TimeString `json:"CreatedAt,omitempty" db:"created_at"`
		UpdatedAt    gorm.TimeString `json:"UpdatedAt,omitempty" db:"updated_at"`
		DeletedAt    gorm.NullString `json:"DeletedAt,omitempty" db:"deleted_at"`
	}
	role struct{}
)

var Role = &role{}

func (*role) GetRole(roleId string) (*AdminRole, error) {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	var role AdminRole
	err := db.Where("nanoid = ?", roleId).First(&role).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &role, err
}
func (*role) GetRoleAll() (roles []*AdminRole, err error) {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	err = db.Find(&roles).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (*role) GetRolesByNames(roleName []string) (roles []*AdminRole, err error) {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	err = db.Where("name IN ?", roleName).Order("id desc").Find(&roles).Error
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
func (*role) GetRoles(offset, limit int) (roles []*AdminRole, count int64, err error) {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	err = db.Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&roles).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (*role) InsertAdminRole(role *AdminRole) (err error) {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	role.Nanoid = idworker.NumberNanoid(16)

	err = db.Create(role).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (*role) UpdateAdminRole(role *AdminRole) (err error) {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	err = db.Updates(role).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (*role) DeleteAdminRole(roleId string) error {
	db := gorm.ClientNew().Model(&AdminRole{}).Where("deleted_at IS NULL")
	deletedAt := gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := db.Where("nanoid = ?", roleId).Updates(AdminRole{DeletedAt: deletedAt}).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (*role) UpdateAdminRoleRules(roleId string, ruleIds []string) error {
	db := gorm.ClientNew().Model(AdminRole{}).Where("deleted_at IS NULL")
	err := db.Where("nanoid = ?", roleId).Update("rules", ruleIds).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}
