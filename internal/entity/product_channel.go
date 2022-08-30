package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/idworker"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	ProductChannel struct {
		gdb.Model
		ID          int    `json:"ID,omitempty" db:"id"`
		ChannelUuid string `json:"ChannelUuid,omitempty" db:"channel_uuid"`

		Title   string `json:"Title,omitempty" db:"title"`
		Channel string `json:"Channel,omitempty" db:"channel"`

		CreatedAt gorm.TimeString `json:"CreatedAt,omitempty" db:"created_at"`
		UpdatedAt gorm.TimeString `json:"UpdatedAt,omitempty" db:"updated_at"`
		DeletedAt gorm.NullString `json:"DeletedAt,omitempty" db:"deleted_at"`
	}
	_productChannel struct{}
)

var ProductChannelEntity = new(_productChannel)

func (*_productChannel) db() *gdb.DB {
	return gorm.NewModel(&ProductChannel{}).Where("deleted_at IS NULL")
}

func (e *_productChannel) ChannelWeb() string {
	return "web.pc"
}

func (e *_productChannel) ChannelMobile() string {
	return "web.mobile"
}

func (e *_productChannel) ChannelApi() string {
	return "api"
}

func (e *_productChannel) GetProductChannel(uuid string) (*ProductChannel, error) {
	var data ProductChannel
	err := e.db().Where("channel_uuid = ?", uuid).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_productChannel) GetProductChannels(offset, limit int) (products []*ProductChannel, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_productChannel) Insert(channel *ProductChannel) (err error) {
	channel.ChannelUuid = idworker.StringNanoid(16)
	err = e.db().Create(channel).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_productChannel) Delete(productId int) error {
	deletedAt := gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := e.db().Where("id = ?", productId).Updates(ProductChannel{DeletedAt: deletedAt}).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_productChannel) GetAll() (products []*ProductChannel, err error) {
	err = e.db().Order("id desc").Find(&products).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
