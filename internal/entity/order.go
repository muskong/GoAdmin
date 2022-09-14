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
		OrderNumber  string       `json:"OrderNumber" db:"order_number"`
		UserId       string       `json:"UserId" db:"user_id"`
		ExternalId   string       `json:"ExternalId" db:"external_id"`
		Queue        string       `json:"Queue" db:"queue"`
		State        string       `json:"State" db:"state"`
		CardNumber   string       `json:"CardNumber" db:"card_number"`
		CardPassword string       `json:"CardPassword" db:"card_password"`
		CardCvv      string       `json:"CardCvv" db:"card_cvv"`
		ResultCard   gorm.JsonAny `json:"ResultCard" db:"result_card"`
		ResultPay    gorm.JsonAny `json:"ResultPay" db:"result_pay"`

		ProductId        int `json:"ProductId,omitempty" db:"product_id"`
		ProductCardId    int `json:"ProductCardId,omitempty" db:"product_card_id"`
		ProductAmountId  int `json:"ProductAmountId,omitempty" db:"product_amount_id"`
		ProductChannelId int `json:"ProductChannelId,omitempty" db:"product_channel_id"`
		ProductServiceId int `json:"ProductServiceId,omitempty" db:"product_service_id"`
		PayId            int `json:"PayId,omitempty" db:"pay_id"`

		Product Product        `gorm:"foreignkey:ID;references:ProductId"`
		Card    ProductCard    `gorm:"foreignkey:ID;references:ProductCardId"`
		Amount  ProductAmount  `gorm:"foreignkey:ID;references:ProductAmountId"`
		Channel ProductChannel `gorm:"foreignkey:ID;references:ProductChannelId"`
		Service ProductService `gorm:"foreignkey:ID;references:ProductServiceId"`
		Pay     Pay            `gorm:"foreignkey:ID;references:PayId"`
	}
	_order struct{}
)

var OrderEntity = new(_order)

func (*_order) db() *gdb.DB {
	return gorm.NewModel(&Order{}).Preload("Product").Preload("Card").Preload("Amount").Preload("Channel").Preload("Service").Preload("Pay").Where("deleted_at IS NULL")
}

func (*_order) ChannelPC() string {
	return "web.pc"
}
func (*_order) ChannelMobile() string {
	return "web.mobile"
}
func (*_order) ChannelApi() string {
	return "api"
}

// 创建订单默认状态
func (*_order) QueueCardQueue() string {
	return "card.queue"
}

// 订单被调用, 发起销卡
func (*_order) QueueCardRequest() string {
	return "card.request"
}

// 销卡成功
func (*_order) QueueCardSuccess() string {
	return "card.success"
}

// 销卡失败
func (*_order) QueueCardError() string {
	return "card.error"
}

// 销卡成功后, 发起付款
func (*_order) QueuePayRequest() string {
	return "pay.request"
}

// 付款成功
func (*_order) QueuePaySuccess() string {
	return "pay.success"
}

// 付款失败
func (*_order) QueuePayError() string {
	return "pay.error"
}

// 订单执行中
func (*_order) StateHang() string {
	return "hang"
}

// 订单执行结束, 成功完成
func (*_order) StateSuccess() string {
	return "success"
}

// 订单失败
func (*_order) StateError() string {
	return "error"
}

func (o *_order) GetOrder(orderId string) (*Order, error) {
	var order Order
	err := o.db().Where("uuid = ?", orderId).First(&order).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &order, err
}

func (e *_order) GetQueueOrders(offset, limit int) (orders []*Order, count int64, err error) {
	err = e.db().Where("state=?", e.StateHang()).Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&orders).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
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
