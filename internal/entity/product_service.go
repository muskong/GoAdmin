package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	ProductService struct {
		gdb.Model
		Id          int    `json:"Id,omitempty" db:"id"`
		ServiceUuid string `json:"ServiceUuid,omitempty" db:"service_uuid"`

		Title       string `json:"Title,omitempty" db:"title"`
		Class       string `json:"Class,omitempty" db:"class"`
		Status      string `json:"Status,omitempty" db:"status"`
		ServiceType string `json:"ServiceType,omitempty" db:"service_type"`

		CreatedAt gorm.TimeString `json:"CreatedAt,omitempty" db:"created_at"`
		UpdatedAt gorm.TimeString `json:"UpdatedAt,omitempty" db:"updated_at"`
		DeletedAt gorm.NullString `json:"DeletedAt,omitempty" db:"deleted_at"`
	}
	_productService struct{}
)

var ProductServiceEntity = new(_productService)

func (*_productService) db() *gdb.DB {
	return gorm.NewModel(&ProductService{}).Where("deleted_at IS NULL")
}

func (e *_productService) StatusAllow() string {
	return "allow"
}
func (e *_productService) StatusDeny() string {
	return "deny"
}

func (e *_productService) GetProductService(uuid string) (*ProductService, error) {
	var data ProductService
	err := e.db().Where("service_uuid = ?", uuid).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_productService) GetProductServices(offset, limit int) (products []*ProductService, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_productService) Insert(product *ProductService) (err error) {
	db := gorm.NewModel(&ProductService{})

	err = db.Create(product).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_productService) Delete(productId int) error {
	db := gorm.NewModel(&ProductService{}).Where("deleted_at IS NULL")
	deletedAt := gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := db.Where("id = ?", productId).Updates(ProductService{DeletedAt: deletedAt}).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_productService) GetAll() (products []*ProductService, err error) {
	err = e.db().Order("id desc").Find(&products).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
