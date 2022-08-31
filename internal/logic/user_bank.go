package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _userBank struct {
	Logic
}

var UserBank = &_userBank{}

func (l *_userBank) List(page Page) (result Result, err error) {

	data, count, err := entity.UserBankEntity.GetUserBanks(page.getOffset(), page.getLimit())
	if len(data) <= 0 || err != nil {
		err = errors.New("用户银行卡无数据")
		return
	}

	result.Data = data
	result.Pagination.Page = page
	result.Pagination.Total = count

	return
}

func (l *_userBank) Create(data entity.UserBank) error {
	err := entity.UserBankEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("新增用户银行卡失败")
	}

	l.Log("新增用户银行卡", data)
	return err
}
func (l *_userBank) Update(data entity.UserBank) (err error) {
	err = entity.UserBankEntity.Delete(data.ID)
	if err != nil {
		return errors.New("更新用户银行卡失败")
	}
	err = entity.UserBankEntity.Insert(&data)
	if data.ID <= 0 || err != nil {
		return errors.New("更新用户银行卡失败")
	}

	l.Log("更新用户银行卡", data)
	return err
}
func (l *_userBank) Delete(userId int) error {
	err := entity.UserBankEntity.Delete(userId)
	if err != nil {
		return errors.New("删除用户银行卡失败")
	}
	l.Log("删除用户银行卡", userId)
	return err
}
