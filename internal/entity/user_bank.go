package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	UserBank struct {
		gorm.Model
		Mobile       string          `json:"Mobile,omitempty" db:"mobile"`
		Province     string          `json:"Province,omitempty" db:"province"`
		City         string          `json:"City,omitempty" db:"city"`
		Type         string          `json:"Type,omitempty" db:"type"`
		Abbreviation string          `json:"Abbreviation,omitempty" db:"abbreviation"`
		Bank         string          `json:"Bank,omitempty" db:"bank"`
		BinDigits    string          `json:"BinDigits,omitempty" db:"bin_digits"`
		CardNumber   string          `json:"CardNumber,omitempty" db:"card_number"`
		CardName     string          `json:"CardName,omitempty" db:"card_name"`
		CardBin      string          `json:"CardBin,omitempty" db:"card_bin"`
		CardDigits   string          `json:"CardDigits,omitempty" db:"card_digits"`
		Logo         string          `json:"Logo,omitempty" db:"logo"`
		IsLuhn       string          `json:"IsLuhn,omitempty" db:"is_luhn"`
		Images       gorm.JsonString `json:"Images,omitempty" db:"images"`
	}
	_userBank struct{}
)

var UserBankEntity = new(_userBank)

func (*_userBank) db() *gdb.DB {
	return gorm.NewModel(&UserBank{})
}

func (e *_userBank) GetUserBank(uuid string) (*UserBank, error) {
	var data UserBank
	err := e.db().Where("uuid = ?", uuid).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_userBank) GetUserBanks(offset, limit int) (userBanks []*UserBank, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&userBanks).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_userBank) Insert(userBank *UserBank) (err error) {
	err = e.db().Create(userBank).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_userBank) Delete(bankId int) error {
	err := e.db().Delete(&UserBank{}, bankId).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_userBank) GetAll() (userBanks []*UserBank, err error) {
	err = e.db().Order("id desc").Find(&userBanks).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
