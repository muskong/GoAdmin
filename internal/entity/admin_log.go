package entity

import (
	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	AdminLogEntity struct {
		/** ID **/
		Id        int             `json:"Id,omitempty" db:"id"`
		AdminId   string          `db:"admin_id"`
		AdminName string          `db:"admin_name"`
		Ip        string          `db:"ip"`
		Url       string          `db:"url"`
		Method    string          `db:"method"`
		Type      string          `db:"type"`
		Param     string          `db:"param"`
		Useragent string          `db:"useragent"`
		Remark    string          `db:"remark"`
		CreatedAt gorm.TimeString `db:"created_at"`
	}
	_log struct{}
)

var AdminLog = new(_log)

func (e *_log) db() *gdb.DB {
	return gorm.NewModel(&OrderEntity{})
}

// 按选项查询集合
func (e *_log) GetAdminLogs(page, limit int) (logs []*AdminLogEntity, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(page).Find(&logs).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_log) InsertAdminLog(log *AdminLogEntity) (*AdminLogEntity, error) {
	err := e.db().Create(log).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return log, err
}
