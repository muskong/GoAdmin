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

func (o *_external) GetExternal(orderId string) (*External, error) {
	var order External
	err := o.db().Where("nanoid = ?", orderId).First(&order).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &order, err
}

func (e *_external) GetExternals(offset, limit int) (orders []*External, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&orders).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_external) Insert(order *External) (err error) {
	// order.ExternalNumber = idworker.IdWorker("C")

	err = e.db().Create(order).Error
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

func (e *_external) GetAll() (orders []*External, err error) {
	err = e.db().Order("id desc").Find(&orders).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
