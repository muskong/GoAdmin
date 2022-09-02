package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/idworker"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	Order struct {
		gorm.Model
		OrderNumber  string `json:"OrderNumber" db:"order_number"`
		ProductId    string `json:"ProductId" db:"product_id"`
		UserId       string `json:"UserId" db:"user_id"`
		ExternalId   string `json:"ExternalId" db:"external_id"`
		Channel      string `json:"Channel" db:"channel"`
		Queue        string `json:"Queue" db:"queue"`
		State        string `json:"State" db:"state"`
		CardNumber   string `json:"CardNumber" db:"card_number"`
		CardPassword string `json:"CardPassword" db:"card_password"`
		CardCvv      string `json:"CardCvv" db:"card_cvv"`
		Result       string `json:"Result" db:"result"`
	}
	_order struct{}
)

var OrderEntity = new(_order)

func (*_order) db() *gdb.DB {
	return gorm.NewModel(&Order{}).Where("deleted_at IS NULL")
}

func (o *_order) GetOrder(orderId string) (*Order, error) {
	var order Order
	err := o.db().Where("nanoid = ?", orderId).First(&order).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &order, err
}

func (e *_order) GetOrders(offset, limit int) (orders []*Order, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&orders).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_order) Insert(order *Order) (err error) {
	order.OrderNumber = idworker.IdWorker("C")

	err = e.db().Create(order).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (e *_order) Delete(orderId int) error {
	err := e.db().Delete(&Order{}, orderId).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_order) GetAll() (orders []*Order, err error) {
	err = e.db().Order("id desc").Find(&orders).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
