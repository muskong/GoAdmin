package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _order struct {
	Logic
}

var Order = &_order{}

func (l *_order) List(page Page) (result Result, err error) {

	data, count, err := entity.OrderEntity.GetOrders(page.getOffset(), page.getLimit())
	if len(data) <= 0 || err != nil {
		err = errors.New("用户认证无数据")
		return
	}

	result.Data = data
	result.Pagination.Page = page
	result.Pagination.Total = count

	return
}

func (l *_order) Create(data entity.Order) error {
	err := entity.OrderEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("新增用户认证失败")
	}

	l.Log("新增用户认证", data)
	return err
}
func (l *_order) Update(data entity.Order) (err error) {
	err = entity.OrderEntity.Delete(data.ID)
	if err != nil {
		return errors.New("更新用户认证失败")
	}
	err = entity.OrderEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("更新用户认证失败")
	}

	l.Log("更新用户认证", data)
	return err
}
func (l *_order) Delete(userId int) error {
	err := entity.OrderEntity.Delete(userId)
	if err != nil {
		return errors.New("删除用户认证失败")
	}
	l.Log("删除用户认证", userId)
	return err
}
