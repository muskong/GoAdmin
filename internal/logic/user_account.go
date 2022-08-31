package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _userAccount struct {
	Logic
}

var UserAccount = &_userAccount{}

func (l *_userAccount) List(page Page) (result Result, err error) {

	data, count, err := entity.UserAccountEntity.GetUserAccounts(page.getOffset(), page.getLimit())
	if len(data) <= 0 || err != nil {
		err = errors.New("用户账户无数据")
		return
	}

	result.Data = data
	result.Pagination.Page = page
	result.Pagination.Total = count

	return
}

func (l *_userAccount) Create(data entity.UserAccount) error {
	err := entity.UserAccountEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("新增用户账户失败")
	}

	l.Log("新增用户账户", data)
	return err
}
