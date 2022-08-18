package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _log struct {
	Logic
}

var Log = &_log{}

func (*_log) AdminLogList(page Page) (result Result, err error) {

	logData, logCount, err := entity.AdminLog.GetAdminLogs(page.getOffset(), page.Limit)
	if len(logData) <= 0 {
		err = errors.New("无数据")
		return
	}

	result.Data = logData
	result.Pagination.Page = page
	result.Pagination.Total = logCount

	return
}
