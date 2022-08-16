package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	OrderEntity struct {
		/** ID **/
		Id     int    `json:"Id,omitempty" db:"id"`
		Nanoid string `json:"Nanoid,omitempty" db:"nanoid"`
		/** 发布时间 **/
		CreatedAt gorm.TimeString `json:"CreatedAt,omitempty" db:"created_at"`
		/** 更新时间 **/
		UpdatedAt gorm.TimeString `json:"UpdatedAt,omitempty" db:"updated_at"`
		DeletedAt gorm.TimeString `json:"DeletedAt,omitempty" db:"deleted_at"`
	}
	_order struct{}
)

var Order = new(_order)

func (*_order) db() *gdb.DB {
	return gorm.NewModel(&OrderEntity{}).Where("deleted_at IS NULL")
}

func (o *_order) GetOrder(orderId string) (*OrderEntity, error) {
	var order OrderEntity
	err := o.db().Where("nanoid = ?", orderId).First(&order).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &order, err
}
