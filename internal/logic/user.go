package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

func AdminUserList(page Page) (err error, result Result) {

	userCount := entity.GetUserCount()
	userData := entity.GetUsers(page.getOffset(), page.Limit, "admin_user_id", "desc")
	if len(userData) <= 0 {
		err = errors.New("无数据")
		return
	}

	result.Data = userData
	result.Pagination.Limit = page.Limit
	result.Pagination.Page = page.Page
	result.Pagination.Total = userCount

	return
}

func AdminUserCreate(u entity.AdminUserObject) error {
	id := entity.InsertAdminUser(u)
	if id <= 0 {
		return errors.New("新增用户失败")
	}
	return nil
}
