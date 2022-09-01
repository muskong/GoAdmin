package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/idworker"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	ProductAmount struct {
		gorm.Model
		AmountUuid string `json:"AmountUuid,omitempty" db:"amount_uuid"`
		Amount     string `json:"Amount,omitempty" db:"amount"`
		Rate       string `json:"Rate,omitempty" db:"rate"`
		RateSys    string `json:"RateSys,omitempty" db:"rate_sys"`
	}
	_productAmount struct{}
)

var ProductAmountEntity = new(_productAmount)

func (*_productAmount) db() *gdb.DB {
	return gorm.NewModel(&ProductAmount{}).Where("deleted_at IS NULL")
}

func (e *_productAmount) GetProductAmount(uuid string) (*ProductAmount, error) {
	var data ProductAmount
	err := e.db().Where("amount_uuid = ?", uuid).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_productAmount) GetProductAmounts(offset, limit int) (products []*ProductAmount, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_productAmount) Insert(product *ProductAmount) (err error) {
	product.AmountUuid = idworker.StringNanoid(16)
	err = e.db().Create(product).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_productAmount) Delete(productId int) error {
	pa := ProductAmount{}
	pa.DeletedAt = gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := e.db().Where("id = ?", productId).Updates(pa).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_productAmount) GetAll() (products []*ProductAmount, err error) {
	err = e.db().Order("id desc").Find(&products).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
