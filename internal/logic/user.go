package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _user struct {
	Logic
}

var User = &_user{}

func (l *_user) List(page Page) (result Result, err error) {

	data, count, err := entity.UserEntity.GetUsers(page.getOffset(), page.getLimit())
	if len(data) <= 0 || err != nil {
		err = errors.New("用户无数据")
		return
	}

	result.Data = data
	result.Pagination.Page = page
	result.Pagination.Total = count

	return
}

func (l *_user) Create(data entity.User) error {
	err := entity.UserEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("新增用户失败")
	}

	l.Log("新增用户", data)
	return err
}
func (l *_user) Update(data entity.User) (err error) {
	err = entity.UserEntity.Delete(data.ID)
	if err != nil {
		return errors.New("更新用户失败")
	}
	err = entity.UserEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("更新用户失败")
	}

	l.Log("更新用户", data)
	return err
}
func (l *_user) Delete(productId int) error {
	err := entity.UserEntity.Delete(productId)
	if err != nil {
		return errors.New("删除用户失败")
	}
	l.Log("删除用户", productId)
	return err
}
