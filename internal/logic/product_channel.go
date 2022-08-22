package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _productChannel struct {
	Logic
}

var ProductChannel = &_productChannel{}

func (l *_productChannel) List(page Page) (result Result, err error) {

	productData, productCount, err := entity.ProductChannelEntity.GetProductChannels(page.getOffset(), page.getLimit())
	if len(productData) <= 0 || err != nil {
		err = errors.New("产品渠道无数据")
		return
	}

	result.Data = productData
	result.Pagination.Page = page
	result.Pagination.Total = productCount

	return
}

func (l *_productChannel) Create(data entity.ProductChannel) error {
	err := entity.ProductChannelEntity.Insert(&data)
	if data.Id <= 0 || err != nil {
		return errors.New("新增产品渠道失败")
	}

	l.Log("新增产品渠道", data)
	return err
}
func (l *_productChannel) Update(data entity.ProductChannel) (err error) {
	err = entity.ProductChannelEntity.Delete(data.Id)
	if err != nil {
		return errors.New("更新产品渠道失败")
	}
	err = entity.ProductChannelEntity.Insert(&data)
	if data.Id <= 0 || err != nil {
		return errors.New("更新产品渠道失败")
	}

	l.Log("更新产品渠道", data)
	return err
}
func (l *_productChannel) Delete(productId int) error {
	err := entity.ProductChannelEntity.Delete(productId)
	if err != nil {
		return errors.New("删除产品渠道失败")
	}
	l.Log("删除产品渠道", productId)
	return err
}
