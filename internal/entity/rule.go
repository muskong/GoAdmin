package entity

import (
	"database/sql"

	"github.com/muskong/GoPkg/zaplog"
)

type (
	AdminRule struct {
		Id        int64          `json:"Id,omitempty" db:"id"`
		Title     string         `json:"Title,omitempty" db:"title"`
		Pid       string         `json:"Pid,omitempty" db:"pid"`
		Link      string         `json:"Link,omitempty" db:"link"`
		Icon      string         `json:"Icon,omitempty" db:"icon"`
		Type      string         `json:"Type,omitempty" db:"type"`
		Active    string         `json:"Active,omitempty" db:"active"`
		Sequence  string         `json:"Sequence,omitempty" db:"sequence"`
		CreatedAt string         `json:"createdAt,omitempty" db:"created_at"`
		UpdatedAt sql.NullString `json:"updatedAt,omitempty" db:"updated_at"`
		DeletedAt sql.NullString `json:"deletedAt,omitempty" db:"deleted_at"`
	}
)

func GetRule(ruleId int64) (*AdminRule, error) {
	var rule AdminRule
	err := db.Where("id = ?", ruleId).First(&rule).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &rule, err
}

func GetRuleAllByPid(pid int64) (rules []*AdminRule, err error) {
	err = db.Where("pid = ?", pid).Find(&rules).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func GetRuleAll() (rules []*AdminRule, err error) {
	err = db.Find(&rules).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

// 按选项查询集合
func GetRules(offset, limit int) (rules []*AdminRule, count int64, err error) {
	err = db.Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&rules).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func InsertAdminRule(rule *AdminRule) (*AdminRule, error) {
	err := db.Create(rule).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return rule, err
}
