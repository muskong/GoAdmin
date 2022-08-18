package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _user struct {
	Logic
}

var User = &_user{}

func (l *_user) AdminUser(userId int) (user *entity.AdminUser, err error) {
	user, err = entity.User.GetUser(userId)
	if user.Id <= 0 || err != nil {
		err = errors.New("查询用户失败")
	}
	return
}
func (l *_user) AdminUserList(page Page) (result Result, err error) {

	userData, userCount, err := entity.User.GetUsers(page.getOffset(), page.Limit)
	if len(userData) <= 0 {
		err = errors.New("无数据")
		return
	}

	result.Data = userData
	result.Pagination.Page = page
	result.Pagination.Total = userCount

	return
}

func (l *_user) AdminUserCreate(u entity.AdminUser) error {
	user, err := entity.User.InsertAdminUser(&u)
	if user.Id <= 0 || err != nil {
		return errors.New("新增用户失败")
	}
	l.Log("新增用户", user)
	return nil
}

func (l *_user) AdminUserUpdate(u entity.AdminUser) error {
	user, err := entity.User.UpdateAdminUser(&u)
	if user.Id <= 0 || err != nil {
		return errors.New("更新用户失败")
	}
	l.Log("更新用户", user)
	return nil
}

func (l *_user) AdminUserDelete(userId int) error {
	err := entity.User.DeleteAdminUser(userId)
	if err != nil {
		return errors.New("删除用户失败")
	}
	l.Log("删除用户", userId)
	return nil
}
