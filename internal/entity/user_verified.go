package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/idworker"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	UserVerified struct {
		gdb.Model
		ID                 int             `json:"ID,omitempty" db:"id"`
		UserUuid           string          `json:"UserUuid,omitempty" db:"user_uuid"`
		Type               string          `json:"Type,omitempty" db:"type"`
		State              string          `json:"State,omitempty" db:"state"`
		Method             string          `json:"Method,omitempty" db:"method"`
		IdentityNumber     string          `json:"IdentityNumber,omitempty" db:"identity_number"`
		UnifiedSocial      string          `json:"UnifiedSocial,omitempty" db:"unified_social"`
		FrontPhoto         string          `json:"FrontPhoto,omitempty" db:"front_photo"`
		BackPhoto          string          `json:"BackPhoto,omitempty" db:"back_photo"`
		HandPhoto          string          `json:"HandPhoto,omitempty" db:"hand_photo"`
		UnifiedSocialPhoto string          `json:"UnifiedSocialPhoto,omitempty" db:"unified_social_photo"`
		CompanyName        string          `json:"CompanyName,omitempty" db:"company_name"`
		Remarks            string          `json:"Remarks,omitempty" db:"remarks"`
		CreatedAt          gorm.TimeString `json:"CreatedAt,omitempty" db:"created_at"`
		UpdatedAt          gorm.TimeString `json:"UpdatedAt,omitempty" db:"updated_at"`
		DeletedAt          gorm.NullString `json:"DeletedAt,omitempty" db:"deleted_at"`
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
