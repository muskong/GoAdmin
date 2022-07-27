package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _user struct{}

var User = &_user{}

func (*_user) AdminUserList(page Page) (err error, result Result) {

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

func (*_user) AdminUserCreate(u entity.AdminUser) error {
	user, err := entity.User.InsertAdminUser(&u)
	if user.Id <= 0 || err != nil {
		return errors.New("新增用户失败")
	}
	return nil
}

func (*_user) GetUser(userId int64) (user *entity.AdminUser, err error) {
	user, err = entity.User.GetUser(userId)
	if user.Id <= 0 || err != nil {
		err = errors.New("新增用户失败")
	}
	return
}
