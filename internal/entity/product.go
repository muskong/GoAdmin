package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	Product struct {
		gorm.Model
		ProductNumber    string `json:"ProductNumber,omitempty" db:"product_number"`
		ProductCardId    int    `json:"ProductCardId,omitempty" db:"product_card_id"`
		ProductAmountId  int    `json:"ProductAmountId,omitempty" db:"product_amount_id"`
		ProductChannelId int    `json:"ProductChannelId,omitempty" db:"product_channel_id"`
		ProductServiceId int    `json:"ProductServiceId,omitempty" db:"product_service_id"`
		Weight           int    `json:"Weight,omitempty" db:"weight"`
		Current          int    `json:"-" db:"-"`
		Status           string `json:"Status,omitempty" db:"status"`

		Card    ProductCard    `gorm:"foreignkey:ID;references:ProductCardId"`
		Amount  ProductAmount  `gorm:"foreignkey:ID;references:ProductAmountId"`
		Channel ProductChannel `gorm:"foreignkey:ID;references:ProductChannelId"`
		Service ProductService `gorm:"foreignkey:ID;references:ProductServiceId"`
	}
	_product struct{}
)

var ProductEntity = new(_product)

func (*_product) db() *gdb.DB {
	return gorm.NewModel(&Product{}).Preload("Card").Preload("Amount").Preload("Channel").Preload("Service").Where("deleted_at IS NULL")
}

func (e *_product) StatusAllow() string {
	return "allow"
}
func (e *_product) StatusDeny() string {
	return "deny"
}

func (e *_product) GetProduct(uuid string) (*Product, error) {
	var data Product
	err := e.db().Where("id = ?", uuid).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_product) GetProducts(offset, limit int) (products []*Product, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_product) GetServices(cardId, amountId, channelId int) (products []*Product, err error) {
	err = e.db().Where("product_card_id=? AND product_amount_id=? AND product_channel_id=?", cardId, amountId, channelId).Find(&products).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_product) Insert(product *Product) (err error) {

	err = e.db().Omit("UpdatedAt").Create(product).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_product) Delete(productId int) error {
	p := Product{}
	p.DeletedAt = gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := e.db().Where("id = ?", productId).Updates(p).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}
