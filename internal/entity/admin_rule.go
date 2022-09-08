package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	rule      struct{}
	AdminRule struct {
		gorm.Model
		ParentUuid string `json:"ParentUuid,omitempty" db:"parent_uuid"`
		Type       string `json:"Type,omitempty" db:"type"`
		Title      string `json:"Title,omitempty" db:"title"`
		Path       string `json:"Path,omitempty" db:"path"`
		Icon       string `json:"Icon,omitempty" db:"icon"`
		Remark     string `json:"Remark,omitempty" db:"remark"`
		Active     string `json:"Active,omitempty" db:"active"`
		Sequence   string `json:"Sequence,omitempty" db:"sequence"`
	}
)

var Rule = &rule{}

func (*rule) db() *gdb.DB {
	return gorm.NewModel(&AdminRule{}).Where("deleted_at IS NULL")
}

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

func (u *rule) GetRule(uuid string) (*AdminRule, error) {

	var rule AdminRule
	err := u.db().Where("uuid = ?", uuid).First(&rule).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &rule, err
}

func (u *rule) GetRuleAllByPid(parentUuid string) (rules []*AdminRule, err error) {

	err = u.db().Where("parent_uuid = ?", parentUuid).Find(&rules).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (u *rule) GetRuleAll() (rules []*AdminRule, err error) {
	err = u.db().Where("active=?", u.ActiveAllow()).Find(&rules).Order("sequence DESC").Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (u *rule) GetRulesByIds(ids []string) (rules []*AdminRule, err error) {
	err = u.db().Where("active=?", u.ActiveAllow()).Where("uuid IN (?)", ids).Order("sequence DESC").Find(&rules).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

// 按选项查询集合
func (u *rule) GetRules(offset, limit int) (rules []*AdminRule, count int64, err error) {
	err = u.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&rules).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (u *rule) InsertAdminRule(rule *AdminRule) (*AdminRule, error) {
	err := u.db().Create(rule).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return rule, err
}
func (u *rule) UpdateAdminRule(rule *AdminRule) (*AdminRule, error) {
	err := u.db().Updates(rule).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return rule, err
}
func (u *rule) DeleteAdminRule(ruleId string) error {
	ar := AdminRule{}
	ar.DeletedAt = gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := u.db().Where("uuid = ?", ruleId).Updates(ar).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}
