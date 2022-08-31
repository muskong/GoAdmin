package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _userGroup struct {
	Logic
}

var UserGroup = &_userGroup{}

func (l *_userGroup) List(page Page) (result Result, err error) {

	data, count, err := entity.UserGroupEntity.GetUserGroups(page.getOffset(), page.getLimit())
	if len(data) <= 0 || err != nil {
		err = errors.New("用户组无数据")
		return
	}

	result.Data = data
	result.Pagination.Page = page
	result.Pagination.Total = count

	return
}

func (l *_userGroup) Create(data entity.UserGroup) error {
	err := entity.UserGroupEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("新增用户组失败")
	}

	l.Log("新增用户组", data)
	return err
}
func (l *_userGroup) Update(data entity.UserGroup) (err error) {
	err = entity.UserGroupEntity.Delete(data.ID)
	if err != nil {
		return errors.New("更新用户组失败")
	}
	err = entity.UserGroupEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("更新用户组失败")
	}

	l.Log("更新用户组", data)
	return err
}
func (l *_userGroup) Delete(userId int) error {
	err := entity.UserGroupEntity.Delete(userId)
	if err != nil {
		return errors.New("删除用户组失败")
	}
	l.Log("删除用户组", userId)
	return err
}
