package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	ProductAmount struct {
		gdb.Model
		Id         int    `json:"Id,omitempty" db:"id"`
		AmountUuid string `json:"AmountUuid,omitempty" db:"amount_uuid"`

		Amount  string `json:"Amount,omitempty" db:"amount"`
		Rate    string `json:"Rate,omitempty" db:"rate"`
		RateSys string `json:"RateSys,omitempty" db:"rate_sys"`

		CreatedAt gorm.TimeString `json:"CreatedAt,omitempty" db:"created_at"`
		UpdatedAt gorm.TimeString `json:"UpdatedAt,omitempty" db:"updated_at"`
		DeletedAt gorm.NullString `json:"DeletedAt,omitempty" db:"deleted_at"`
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
	db := gorm.NewModel(&ProductAmount{})

	err = db.Create(product).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_productAmount) Delete(productId int) error {
	db := gorm.NewModel(&ProductAmount{}).Where("deleted_at IS NULL")
	deletedAt := gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := db.Where("id = ?", productId).Updates(ProductAmount{DeletedAt: deletedAt}).Error
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
