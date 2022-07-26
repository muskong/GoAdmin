package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

func AdminRoleList(page Page) (err error, result Result) {

	userData, userCount, err := entity.GetRoles(page.getOffset(), page.Limit)
	if len(userData) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	result.Data = userData
	result.Pagination.Page = page
	result.Pagination.Total = userCount

	return
}

func AdminRoleAll() (err error, result []*entity.AdminRole) {

	result, err = entity.GetRoleAll()
	if len(result) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	return
}

func AdminRoleRuleList(roleId int64) (err error, result []SelectInterface) {

	ruleData, err := entity.GetRuleAll()
	if len(ruleData) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	roleRuleSelected := map[int64]bool{}
	rule, err := entity.GetRole(roleId)
	if rule.Id > 0 && err == nil {
		for _, v := range rule.Rules {
			roleRuleSelected[v] = true
		}
	}

	for _, rule := range ruleData {
		ruleSelect := SelectInterface{
			rule.Id,
			rule.Title,
			false}
		if roleRuleSelected[rule.Id] {
			ruleSelect.Checked = true
		}
		result = append(result, ruleSelect)
	}

	return
}

func AdminRoleCreate(u entity.AdminRole) error {
	err := entity.InsertAdminRole(&u)
	if u.Id <= 0 || err != nil {
		return errors.New("新增失败")
	}
	return err
}

func AdminRoleSaveRule(ou RoleRuleObject) error {
	err := entity.UpdateAdminRoleRules(ou.RoleId, ou.RuleIds)
	if err != nil {
		return errors.New("设置失败")
	}
	return err
}
