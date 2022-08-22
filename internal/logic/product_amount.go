package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _productAmount struct {
	Logic
}

var ProductAmount = &_productAmount{}

func (l *_productAmount) List(page Page) (result Result, err error) {

	productData, productCount, err := entity.ProductAmountEntity.GetProductAmounts(page.getOffset(), page.getLimit())
	if len(productData) <= 0 || err != nil {
		err = errors.New("产品金额无数据")
		return
	}

	result.Data = productData
	result.Pagination.Page = page
	result.Pagination.Total = productCount

	return
}

func (l *_productAmount) Create(data entity.ProductAmount) error {
	err := entity.ProductAmountEntity.Insert(&data)
	if data.Id <= 0 || err != nil {
		return errors.New("新增产品金额失败")
	}

	l.Log("新增产品金额", data)
	return err
}
func (l *_productAmount) Update(data entity.ProductAmount) (err error) {
	err = entity.ProductAmountEntity.Delete(data.Id)
	if err != nil {
		return errors.New("更新产品金额失败")
	}
	err = entity.ProductAmountEntity.Insert(&data)
	if data.Id <= 0 || err != nil {
		return errors.New("更新产品金额失败")
	}

	l.Log("更新产品金额", data)
	return err
}
func (l *_productAmount) Delete(productId int) error {
	err := entity.ProductAmountEntity.Delete(productId)
	if err != nil {
		return errors.New("删除产品金额失败")
	}
	l.Log("删除产品金额", productId)
	return err
}
