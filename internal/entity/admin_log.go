package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	AdminLogs struct {
		gorm.Model
		AdminId   string `db:"admin_id"`
		Ip        string `db:"ip"`
		Url       string `db:"url"`
		Method    string `db:"method"`
		Type      string `db:"type"`
		Param     string `db:"param"`
		Useragent string `db:"useragent"`
		Title     string `db:"title"`
	}
	_log struct{}
)

var AdminLogEntity = new(_log)

func (e *_log) db() *gdb.DB {
	return gorm.NewModel(&AdminLogs{}).Omit("UpdatedAt", "DeletedAt")
}

// 按选项查询集合
func (e *_log) GetAdminLogs(page, limit int) (logs []*AdminLogs, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(page).Find(&logs).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_log) InsertAdminLog(log *AdminLogs) (*AdminLogs, error) {
	err := e.db().Create(log).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return log, err
}
