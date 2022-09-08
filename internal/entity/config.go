package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	Config struct {
		gorm.Model
		Type  string `json:"Type" db:"type"`
		Title string `json:"Title" db:"title"`
		Name  string `json:"Name" db:"name"`
		Value string `json:"Value" db:"value"`
	}
	_config struct{}
)

var ConfigEntity = new(_config)

func (*_config) db() *gdb.DB {
	return gorm.NewModel(&Config{}).Where("deleted_at IS NULL")
}

func (o *_config) GetConfig(configId string) (*Config, error) {
	var config Config
	err := o.db().Where("uuid = ?", configId).First(&config).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &config, err
}

func (e *_config) GetConfigs(offset, limit int) (configs []*Config, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&configs).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_config) Insert(config *Config) (err error) {
	err = e.db().Create(config).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
func (e *_config) Delete(configId int) error {
	err := e.db().Delete(&Config{}, configId).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_config) GetAll() (configs []*Config, err error) {
	err = e.db().Order("id desc").Find(&configs).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
