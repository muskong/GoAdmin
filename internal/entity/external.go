package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	External struct {
		gorm.Model
		Type      string `json:"Type" db:"type"`
		Title     string `json:"Title" db:"title"`
		AppKey    string `json:"AppKey" db:"app_key"`
		AppSecret string `json:"AppSecret" db:"app_secret"`
		Status    string `json:"Status" db:"status"`
	}
	_external struct{}
)

var ExternalEntity = new(_external)

func (*_external) db() *gdb.DB {
	return gorm.NewModel(&External{}).Where("deleted_at IS NULL")
}

func (o *_external) GetExternal(uuid string) (*External, error) {
	var external External
	err := o.db().Where("uuid = ?", uuid).First(&external).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &external, err
}

func (e *_external) GetExternals(offset, limit int) (externals []*External, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&externals).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_external) Insert(external *External) (err error) {
	err = e.db().Create(external).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (e *_external) Delete(verifiedId int) error {
	err := e.db().Delete(&External{}, verifiedId).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_external) GetAll() (externals []*External, err error) {
	err = e.db().Order("id desc").Find(&externals).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
