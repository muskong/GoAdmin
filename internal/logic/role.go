package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

func AdminRoleList(page Page) (err error, result Result) {

	userCount := entity.GetRoleCount()
	userData := entity.GetRoles(page.getOffset(), page.Limit, "admin_role_id", "desc")
	if len(userData) <= 0 {
		err = errors.New("无数据")
		return
	}

	result.Data = userData
	result.Pagination.Limit = page.Limit
	result.Pagination.Page = page.Page
	result.Pagination.Total = userCount

	return
}

func AdminRoleAll() (err error, result []entity.AdminRoleObject) {

	result = entity.GetRoleAll()
	if len(result) <= 0 {
		err = errors.New("无数据")
		return
	}

	return
}

func AdminRoleRuleList(roleId int64) (err error, result []SelectInterface) {

	ruleData := entity.GetRuleAll()
	if len(ruleData) <= 0 {
		err = errors.New("无数据")
		return
	}

	roleRuleData := entity.GetRuleAllByRoleId(roleId)
	if len(roleRuleData) <= 0 {
		err = errors.New("无数据")
		return
	}
	roleRuleSelected := map[int64]bool{}
	for _, v := range roleRuleData {
		roleRuleSelected[v.RuleId] = true
	}

	for _, rule := range ruleData {
		ruleSelect := SelectInterface{
			rule.AdminRuleId,
			rule.AdminRuleTitle,
			false}
		if roleRuleSelected[rule.AdminRuleId] {
			ruleSelect.Checked = true
		}
		result = append(result, ruleSelect)
	}

	return
}

func AdminRoleCreate(u entity.AdminRoleObject) error {
	id := entity.InsertAdminRole(u)
	if id <= 0 {
		return errors.New("新增失败")
	}
	return nil
}

func AdminRoleSaveRule(roleRule RoleRuleObject) error {
	var row int64 = 0
	data := entity.AdminRoleRuleObject{
		RoleId: roleRule.RoleId,
		RuleId: roleRule.RuleId,
	}
	if roleRule.Checked {
		row = entity.InsertAdminRoleRule(data)
	} else {
		row = entity.DeleteAdminRoleRule(data)
	}
	if row <= 0 {
		return errors.New("设置失败")
	}
	return nil
}
