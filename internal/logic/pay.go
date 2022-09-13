package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _pay struct {
	Logic
}

var Pay = &_pay{}

func (l *_pay) List(page Page) (result Result, err error) {

	productData, productCount, err := entity.PayEntity.GetPays(page.getOffset(), page.getLimit())
	if len(productData) <= 0 || err != nil {
		err = errors.New("产品无数据")
		return
	}

	result.Data = productData
	result.Pagination.Page = page
	result.Pagination.Total = productCount

	return
}

func (l *_pay) Create(data entity.Pay) error {
	err := entity.PayEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("新增产品失败")
	}

	l.Log("新增产品", data)
	return err
}
func (l *_pay) Update(data entity.Pay) (err error) {
	err = entity.PayEntity.Delete(data.ID)
	if err != nil {
		return errors.New("更新产品失败")
	}
	err = entity.PayEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("更新产品失败")
	}
	l.Log("更新产品", data)
	return err
}
func (l *_pay) Delete(productId int) error {
	err := entity.PayEntity.Delete(productId)
	if err != nil {
		return errors.New("删除产品失败")
	}
	l.Log("删除产品", productId)
	return err
}

var weightPayNodes []*entity.Pay

func (l *_pay) WeightService() (service *entity.Pay, err error) {

	if len(weightPayNodes) == 0 {
		weightPayNodes, err = entity.PayEntity.GetAll()
		if len(weightPayNodes) <= 0 || err != nil {
			err = errors.New("产品无数据")
			return
		}

	}

	pay := smoothPayWeight(weightPayNodes)

	service, err = entity.PayEntity.GetPayById(pay.ID)
	// if len(productData) <= 0 || err != nil {
	// 	err = errors.New("产品无数据")
	// 	return
	// }

	return
}
func smoothPayWeight(nodes []*entity.Pay) (best *entity.Pay) {
	if len(nodes) == 0 {
		return
	}

	total := 0
	for _, node := range nodes {
		if node == nil {
			continue
		}

		total += node.Weight
		node.Current += node.Weight

		if best == nil || node.Current > best.Current {
			best = node
		}
	}

	if best == nil {
		return
	}

	best.Current -= total

	return
}
