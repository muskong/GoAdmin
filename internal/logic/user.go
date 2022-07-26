package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

func AdminUserList(page Page) (err error, result Result) {

	userData, userCount, err := entity.GetUsers(page.getOffset(), page.Limit)
	if len(userData) <= 0 {
		err = errors.New("无数据")
		return
	}

	result.Data = userData
	result.Pagination.Page = page
	result.Pagination.Total = userCount

	return
}

func AdminUserCreate(u entity.AdminUser) error {
	user, err := entity.InsertAdminUser(&u)
	if user.Id <= 0 || err != nil {
		return errors.New("新增用户失败")
	}
	return nil
}
