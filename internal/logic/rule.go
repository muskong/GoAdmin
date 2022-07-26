package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

func AdminRuleList(page Page) (err error, result Result) {

	ruleCount := entity.GetRuleCount()
	ruleData := entity.GetRules(page.getOffset(), page.Limit, "admin_rule_id", "desc")
	if len(ruleData) <= 0 {
		err = errors.New("无数据")
		return
	}

	result.Data = ruleData
	result.Pagination.Limit = page.Limit
	result.Pagination.Page = page.Page
	result.Pagination.Total = ruleCount

	return
}

func AdminRuleCreate(u entity.AdminRuleObject) error {
	id := entity.InsertAdminRule(u)
	if id <= 0 {
		return errors.New("新增失败")
	}
	return nil
}

func AdminRuleAll(id int64) (err error, result Result) {

	ruleData := entity.GetRuleAllByPid(id)
	if len(ruleData) <= 0 {
		err = errors.New("无数据")
		return
	}
	ruleAll := []SelectInterface{}
	for _, rule := range ruleData {
		ruleAll = append(ruleAll, SelectInterface{
			rule.AdminRuleId,
			rule.AdminRuleTitle,
			false})
	}
	result.Data = ruleAll

	return
}
