package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

func AdminRuleList(page Page) (err error, result Result) {

	ruleData, ruleCount, err := entity.GetRules(page.getOffset(), page.Limit)
	if len(ruleData) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	result.Data = ruleData
	result.Pagination.Page = page
	result.Pagination.Total = ruleCount

	return
}

func AdminRuleCreate(u entity.AdminRule) error {
	rule, err := entity.InsertAdminRule(&u)
	if rule.Id <= 0 || err != nil {
		return errors.New("新增失败")
	}
	return nil
}

func AdminRuleAll(id int64) (err error, result Result) {

	ruleData, err := entity.GetRuleAllByPid(id)
	if len(ruleData) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}
	ruleAll := []SelectInterface{}
	for _, rule := range ruleData {
		ruleAll = append(ruleAll, SelectInterface{
			rule.Id,
			rule.Title,
			false})
	}
	result.Data = ruleAll

	return
}
