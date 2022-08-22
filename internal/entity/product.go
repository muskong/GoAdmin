package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	Product struct {
		gdb.Model
		Id int `json:"Id,omitempty" db:"id"`

		ProductCardId    string `json:"ProductCardId,omitempty" db:"product_card_id"`
		ProductAmountId  string `json:"ProductAmountId,omitempty" db:"product_amount_id"`
		ProductChannelId string `json:"ProductChannelId,omitempty" db:"product_channel_id"`
		ProductServiceId string `json:"ProductServiceId,omitempty" db:"product_service_id"`
		Weight           string `json:"Weight,omitempty" db:"weight"`
		Status           string `json:"Status,omitempty" db:"status"`

		CreatedAt gorm.TimeString `json:"CreatedAt,omitempty" db:"created_at"`
		DeletedAt gorm.NullString `json:"DeletedAt,omitempty" db:"deleted_at"`

		Card    ProductCard    `gorm:"foreignkey:Id;references:ProductCardId"`
		Amount  ProductAmount  `gorm:"foreignkey:Id;references:ProductAmountId"`
		Channel ProductChannel `gorm:"foreignkey:Id;references:ProductChannelId"`
		Service ProductService `gorm:"foreignkey:Id;references:ProductServiceId"`
	}
	_product struct{}
)

var ProductEntity = new(_product)

func (*_product) db() *gdb.DB {
	return gorm.NewModel(&Product{}).Preload("Card", "Amount", "Channel", "Service").Where("deleted_at IS NULL")
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

func (e *_product) Insert(product *Product) (err error) {
	db := gorm.NewModel(&Product{})

	err = db.Create(product).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_product) Delete(productId int) error {
	db := gorm.NewModel(&Product{}).Where("deleted_at IS NULL")
	deletedAt := gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := db.Where("id = ?", productId).Updates(Product{DeletedAt: deletedAt}).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}
