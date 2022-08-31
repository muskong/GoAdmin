package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _userVerified struct {
	Logic
}

var UserVerified = &_userVerified{}

func (l *_userVerified) List(page Page) (result Result, err error) {

	data, count, err := entity.UserVerifiedEntity.GetUserVerifieds(page.getOffset(), page.getLimit())
	if len(data) <= 0 || err != nil {
		err = errors.New("用户认证无数据")
		return
	}

	result.Data = data
	result.Pagination.Page = page
	result.Pagination.Total = count

	return
}

func (l *_userVerified) Create(data entity.UserVerified) error {
	err := entity.UserVerifiedEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("新增用户认证失败")
	}

	l.Log("新增用户认证", data)
	return err
}
func (l *_userVerified) Update(data entity.UserVerified) (err error) {
	err = entity.UserVerifiedEntity.Delete(data.ID)
	if err != nil {
		return errors.New("更新用户认证失败")
	}
	err = entity.UserVerifiedEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("更新用户认证失败")
	}

	l.Log("更新用户认证", data)
	return err
}
func (l *_userVerified) Delete(userId int) error {
	err := entity.UserVerifiedEntity.Delete(userId)
	if err != nil {
		return errors.New("删除用户认证失败")
	}
	l.Log("删除用户认证", userId)
	return err
}
