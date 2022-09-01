package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/idworker"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	UserVerified struct {
		gorm.Model
		UserUuid   string `json:"UserUuid" db:"user_uuid"`
		Type       string `json:"Type" db:"type"`
		State      string `json:"State" db:"state"`
		Method     string `json:"Method" db:"method"`
		Name       string `json:"Name" db:"name"`
		Number     string `json:"IdentityNumber" db:"number"`
		FrontPhoto string `json:"FrontPhoto" db:"front_photo"`
		BackPhoto  string `json:"BackPhoto" db:"back_photo"`
		HandPhoto  string `json:"HandPhoto" db:"hand_photo"`
		Address    string `json:"Address" db:"address"`
		Remarks    string `json:"Remarks" db:"remarks"`
	}
	_userVerified struct{}
)

var UserVerifiedEntity = new(_userVerified)

func (*_userVerified) db() *gdb.DB {
	return gorm.NewModel(&UserVerified{})
}

func (e *_userVerified) GetUserVerified(uuid string) (*UserVerified, error) {
	var data UserVerified
	err := e.db().Where("uuid = ?", uuid).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_userVerified) GetUserVerifieds(offset, limit int) (userVerifieds []*UserVerified, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&userVerifieds).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_userVerified) Insert(userVerified *UserVerified) (err error) {
	userVerified.UserUuid = idworker.StringNanoid(30)

	err = e.db().Create(userVerified).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (e *_userVerified) Delete(verifiedId int) error {
	err := e.db().Delete(&UserVerified{}, verifiedId).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_userVerified) GetAll() (userVerifieds []*UserVerified, err error) {
	err = e.db().Order("id desc").Find(&userVerifieds).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
