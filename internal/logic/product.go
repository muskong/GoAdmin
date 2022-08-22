package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _product struct {
	Logic
}

var Product = &_product{}

func (l *_product) Selects() map[string]any {
	var data map[string]any

	card, err := entity.ProductCardEntity.GetAll()
	if err == nil {
		data["cards"] = card
	}
	amount, err := entity.ProductAmountEntity.GetAll()
	if err == nil {
		data["amounts"] = amount
	}
	channel, err := entity.ProductChannelEntity.GetAll()
	if err == nil {
		data["channels"] = channel
	}
	service, err := entity.ProductServiceEntity.GetAll()
	if err == nil {
		data["services"] = service
	}

	return data
}

func (l *_product) List(page Page) (result Result, err error) {

	productData, productCount, err := entity.ProductEntity.GetProducts(page.getOffset(), page.getLimit())
	if len(productData) <= 0 || err != nil {
		err = errors.New("产品无数据")
		return
	}

	result.Data = productData
	result.Pagination.Page = page
	result.Pagination.Total = productCount

	return
}

func (l *_product) Create(data entity.Product) error {
	err := entity.ProductEntity.Insert(&data)
	if data.Id <= 0 || err != nil {
		return errors.New("新增产品失败")
	}

	l.Log("新增产品", data)
	return err
}
func (l *_product) Update(data entity.Product) (err error) {
	err = entity.ProductEntity.Delete(data.Id)
	if err != nil {
		return errors.New("更新产品失败")
	}
	err = entity.ProductEntity.Insert(&data)
	if data.Id <= 0 || err != nil {
		return errors.New("更新产品失败")
	}
	l.Log("更新产品", data)
	return err
}
func (l *_product) Delete(productId int) error {
	err := entity.ProductEntity.Delete(productId)
	if err != nil {
		return errors.New("删除产品失败")
	}
	l.Log("删除产品", productId)
	return err
}
