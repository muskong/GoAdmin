package entity

import (
	"database/sql"
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
)

type (
	rule      struct{}
	AdminRule struct {
		Id        int            `json:"Id,omitempty" db:"id"`
		Pid       int            `json:"Pid,omitempty" db:"pid"`
		Type      string         `json:"Type,omitempty" db:"type"`
		Title     string         `json:"Title,omitempty" db:"title"`
		Link      string         `json:"Link,omitempty" db:"link"`
		Path      string         `json:"Path,omitempty" db:"path"`
		Icon      string         `json:"Icon,omitempty" db:"icon"`
		MenuType  string         `json:"MenuType,omitempty" db:"menu_type"`
		Url       string         `json:"Url,omitempty" db:"url"`
		Component string         `json:"Component,omitempty" db:"component"`
		Keepalive string         `json:"Keepalive,omitempty" db:"keepalive"`
		Extend    string         `json:"Extend,omitempty" db:"extend"`
		Remark    string         `json:"Remark,omitempty" db:"remark"`
		Active    string         `json:"Active,omitempty" db:"active"`
		Sequence  string         `json:"Sequence,omitempty" db:"sequence"`
		CreatedAt sql.NullString `json:"createdAt,omitempty" db:"created_at"`
		UpdatedAt sql.NullString `json:"updatedAt,omitempty" db:"updated_at"`
		DeletedAt sql.NullString `json:"deletedAt,omitempty" db:"deleted_at"`
	}
)

var Rule = &rule{}

func (m *rule) KeepaliveAllow() string {
	return "allow"
}
func (m *rule) KeepaliveDeny() string {
	return "deny"
}

func (m *rule) ActiveAllow() string {
	return "allow"
}
func (m *rule) ActiveDeny() string {
	return "deny"
}

func (u *rule) GetRule(ruleId int) (*AdminRule, error) {
	db := gorm.ClientNew().Model(AdminRule{}).Where("deleted_at IS NULL")
	var rule AdminRule
	err := db.Where("id = ?", ruleId).First(&rule).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &rule, err
}

func (u *rule) GetRuleAllByPid(pid int) (rules []*AdminRule, err error) {
	db := gorm.ClientNew().Model(AdminRule{}).Where("deleted_at IS NULL")
	err = db.Where("pid = ?", pid).Find(&rules).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (u *rule) GetRuleAll() (rules []*AdminRule, err error) {
	db := gorm.ClientNew().Model(AdminRule{}).Where("deleted_at IS NULL")
	err = db.Where("active=?", u.ActiveAllow()).Find(&rules).Order("pid ASC").Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (u *rule) GetRulesByIds(ids []int) (rules []*AdminRule, err error) {
	db := gorm.ClientNew().Model(AdminRule{}).Where("deleted_at IS NULL")
	err = db.Where("active=?", u.ActiveAllow()).Where("id IN (?)", ids).Order("pid ASC").Find(&rules).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

// 按选项查询集合
func (u *rule) GetRules(offset, limit int) (rules []*AdminRule, count int64, err error) {
	db := gorm.ClientNew().Model(AdminRule{}).Where("deleted_at IS NULL")
	err = db.Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&rules).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (u *rule) InsertAdminRule(rule *AdminRule) (*AdminRule, error) {
	db := gorm.ClientNew().Model(AdminRule{}).Where("deleted_at IS NULL")
	err := db.Create(rule).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return rule, err
}
func (u *rule) UpdateAdminRule(rule *AdminRule) (*AdminRule, error) {
	db := gorm.ClientNew().Model(AdminRule{}).Where("deleted_at IS NULL")
	err := db.Updates(rule).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return rule, err
}
func (u *rule) DeleteAdminRule(ruleId int) error {
	db := gorm.ClientNew().Model(&AdminRule{}).Where("deleted_at IS NULL")
	deletedAt := sql.NullString{
		String: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := db.Where("id=?", ruleId).Updates(AdminRule{DeletedAt: deletedAt}).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}
