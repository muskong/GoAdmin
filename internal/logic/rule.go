package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _rule struct{}

var Rule = &_rule{}

func (*_rule) AdminRule(ruleId int) (rule *entity.AdminRule, err error) {

	rule, err = entity.Rule.GetRule(ruleId)
	if rule.Id <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	return
}

func (*_rule) AdminRuleList(page Page) (err error, result Result) {

	ruleData, ruleCount, err := entity.Rule.GetRules(page.getOffset(), page.Limit)
	if len(ruleData) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	result.Data = ruleData
	result.Pagination.Page = page
	result.Pagination.Total = ruleCount

	return
}

func (*_rule) AdminRuleCreate(u entity.AdminRule) error {
	rule, err := entity.Rule.InsertAdminRule(&u)
	if rule.Id <= 0 || err != nil {
		return errors.New("新增失败")
	}
	return err
}

func (*_rule) AdminRuleUpdate(u entity.AdminRule) error {
	rule, err := entity.Rule.UpdateAdminRule(&u)
	if rule.Id <= 0 || err != nil {
		return errors.New("更新失败")
	}
	return err
}

func (*_rule) AdminRuleDelete(ruleId int) error {
	err := entity.Rule.DeleteAdminRule(ruleId)
	if err != nil {
		return errors.New("删除失败")
	}
	return err
}

func (*_rule) AdminRuleAll(id int) (err error, result Result) {

	ruleData, err := entity.Rule.GetRuleAllByPid(id)
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

func (*_rule) AdminRuleGroup(roles []string) (result any, err error) {
	roleData, err := entity.Role.GetRolesByNames(roles)
	if len(roleData) <= 0 || err != nil {
		err = errors.New("无角色数据")
		return
	}

	var ids []int
	for _, v := range roleData {
		ids = append(ids, v.Rules...)
	}
	var isAll bool
	for _, id := range ids {
		if id == 0 {
			isAll = true
		}
	}
	var ruleData []*entity.AdminRule
	if isAll {
		ruleData, err = entity.Rule.GetRuleAll()
	} else {
		ruleData, err = entity.Rule.GetRulesByIds(ids)
	}

	if len(ruleData) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	result = _tree(ruleData, 0)

	return
}

func _tree(rules []*entity.AdminRule, pid int) []RuleTree {

	pdata := map[int][]RuleTreeNode{}
	for _, rule := range rules {
		pdata[rule.Pid] = append(pdata[rule.Pid], RuleTreeNode{
			Id:        rule.Id,
			Pid:       rule.Pid,
			Type:      rule.Type,
			Title:     rule.Title,
			Link:      rule.Link,
			Path:      rule.Path,
			Icon:      rule.Icon,
			MenuType:  rule.MenuType,
			Url:       rule.Url,
			Component: rule.Component,
			Extend:    rule.Extend,
			KeepAlive: rule.Keepalive,
		})
	}

	return _ruleChildren(pid, pdata)
}

func _ruleChildren(pid int, pdata map[int][]RuleTreeNode) (tree []RuleTree) {
	for _, rule := range pdata[pid] {
		var _tree RuleTree
		_tree.RuleTreeNode = rule
		_tree.Children = _ruleChildren(rule.Id, pdata)

		tree = append(tree, _tree)
	}
	return tree
}

func (*_rule) AdminRuleTree(roles []string) (result any, err error) {
	roleData, err := entity.Role.GetRolesByNames(roles)
	if len(roleData) <= 0 || err != nil {
		err = errors.New("无角色数据")
		return
	}

	var ids []int
	for _, v := range roleData {
		ids = append(ids, v.Rules...)
	}
	var isAll bool
	for _, id := range ids {
		if id == 0 {
			isAll = true
		}
	}
	var ruleData []*entity.AdminRule
	if isAll {
		ruleData, err = entity.Rule.GetRuleAll()
	} else {
		ruleData, err = entity.Rule.GetRulesByIds(ids)
	}

	if len(ruleData) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	pdata := map[string][]AntTreeNode{}

	for _, item := range ruleData {
		pdata[item.ParentNanoid] = append(pdata[item.ParentNanoid], AntTreeNode{
			Value: item.Nanoid,
			Title: item.Title,
		})
	}

	result = AntdTree("", pdata)

	return
}
