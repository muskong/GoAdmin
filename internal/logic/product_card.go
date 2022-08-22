package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _productCard struct {
	Logic
}

var ProductCard = &_productCard{}

func (l *_productCard) List(page Page) (result Result, err error) {

	productData, productCount, err := entity.ProductCardEntity.GetProductCards(page.getOffset(), page.getLimit())
	if len(productData) <= 0 || err != nil {
		err = errors.New("产品卡类无数据")
		return
	}

	result.Data = productData
	result.Pagination.Page = page
	result.Pagination.Total = productCount

	return
}

func (l *_productCard) Create(data entity.ProductCard) error {
	err := entity.ProductCardEntity.Insert(&data)
	if data.Id <= 0 || err != nil {
		return errors.New("新增产品卡类失败")
	}

	l.Log("新增产品卡类", data)
	return err
}
func (l *_productCard) Update(data entity.ProductCard) (err error) {
	err = entity.ProductCardEntity.Delete(data.Id)
	if err != nil {
		return errors.New("更新产品卡类失败")
	}
	err = entity.ProductCardEntity.Insert(&data)
	if data.Id <= 0 || err != nil {
		return errors.New("更新产品卡类失败")
	}

	l.Log("更新产品卡类", data)
	return err
}
func (l *_productCard) Delete(productId int) error {
	err := entity.ProductCardEntity.Delete(productId)
	if err != nil {
		return errors.New("删除产品卡类失败")
	}
	l.Log("删除产品卡类", productId)
	return err
}
