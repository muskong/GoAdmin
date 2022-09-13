package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	Pay struct {
		gorm.Model
		Title   string             `json:"Title,omitempty" db:"title"`
		Class   string             `json:"Class,omitempty" db:"class"`
		Status  string             `json:"Status,omitempty" db:"status"`
		Content gorm.JsonMapString `json:"Content,omitempty" db:"content"`
		Type    string             `json:"Type,omitempty" db:"type"`
		Weight  int                `json:"Weight,omitempty" db:"weight"`
		Current int                `json:"-" db:"-"`
	}
	_pay struct{}
)

var PayEntity = new(_pay)

func (*_pay) db() *gdb.DB {
	return gorm.NewModel(&Pay{}).Where("deleted_at IS NULL")
}

func (e *_pay) StatusAllow() string {
	return "allow"
}
func (e *_pay) StatusDeny() string {
	return "deny"
}

func (e *_pay) GetPay(uuid string) (*Pay, error) {
	var data Pay
	err := e.db().Where("uuid = ?", uuid).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_pay) GetPayById(id int) (*Pay, error) {
	var data Pay
	err := e.db().Where("id = ?", id).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_pay) GetPays(offset, limit int) (services []*Pay, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&services).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_pay) Insert(service *Pay) (err error) {
	err = e.db().Create(service).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_pay) Delete(payId int) error {
	ps := Pay{}
	ps.DeletedAt = gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := e.db().Where("id = ?", payId).Updates(ps).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_pay) GetAll() (services []*Pay, err error) {
	err = e.db().Order("id desc").Find(&services).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
