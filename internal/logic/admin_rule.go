package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _rule struct {
	Logic
}

var Rule = &_rule{}

func (l *_rule) AdminRule(ruleId string) (rule *entity.AdminRule, err error) {

	rule, err = entity.Rule.GetRule(ruleId)
	if rule.Id <= 0 || err != nil {
		err = errors.New("权限节点无数据")
		return
	}

	return
}

func (l *_rule) AdminRuleList(page Page) (result Result, err error) {

	ruleData, ruleCount, err := entity.Rule.GetRules(page.getOffset(), page.Limit)
	if len(ruleData) <= 0 || err != nil {
		err = errors.New("权限节点无数据")
		return
	}

	result.Data = ruleData
	result.Pagination.Page = page
	result.Pagination.Total = ruleCount

	return
}

func (l *_rule) AdminRuleCreate(u entity.AdminRule) error {
	rule, err := entity.Rule.InsertAdminRule(&u)
	if rule.Id <= 0 || err != nil {
		return errors.New("新增权限节点失败")
	}
	l.Log("新增权限节点", rule)
	return err
}

func (l *_rule) AdminRuleUpdate(u entity.AdminRule) error {
	rule, err := entity.Rule.UpdateAdminRule(&u)
	if rule.Id <= 0 || err != nil {
		return errors.New("更新权限节点失败")
	}
	l.Log("更新权限节点", rule)
	return err
}

func (l *_rule) AdminRuleDelete(ruleId string) error {
	err := entity.Rule.DeleteAdminRule(ruleId)
	if err != nil {
		return errors.New("删除权限节点失败")
	}
	l.Log("删除权限节点", ruleId)
	return err
}

func (l *_rule) AdminRuleAll(id string) (result Result, err error) {

	ruleData, err := entity.Rule.GetRuleAllByPid(id)
	if len(ruleData) <= 0 || err != nil {
		err = errors.New("权限节点无数据")
		return
	}
	ruleAll := []SelectInterface{}
	for _, rule := range ruleData {
		ruleAll = append(ruleAll, SelectInterface{
			rule.Nanoid,
			rule.Title,
			false})
	}
	result.Data = ruleAll

	return
}

func (l *_rule) AdminRuleGroup(roles []string) (result any, err error) {
	roleData, err := entity.Role.GetRolesByNames(roles)
	if len(roleData) <= 0 || err != nil {
		err = errors.New("无角色数据")
		return
	}

	var ids []string
	for _, v := range roleData {
		ids = append(ids, v.Rules...)
	}
	var isAll bool
	for _, id := range ids {
		if id == "" {
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
		err = errors.New("权限节点无数据")
		return
	}

	result = _tree(ruleData, 0)

	return
}

func (l *_rule) AdminRuleGroupList() (result any, err error) {
	var ruleData []*entity.AdminRule
	ruleData, err = entity.Rule.GetRuleAll()
	if len(ruleData) <= 0 || err != nil {
		err = errors.New("权限节点无数据")
		return
	}

	result = _tree(ruleData, 0)

	return
}

func _tree(ruleData []*entity.AdminRule, pid int) []RuleTree {

	pdata := map[string][]RuleTreeNode{}
	for _, rule := range ruleData {
		pdata[rule.ParentNanoid] = append(pdata[rule.ParentNanoid], RuleTreeNode{
			Id:           rule.Id,
			Nanoid:       rule.Nanoid,
			ParentNanoid: rule.ParentNanoid,
			Type:         rule.Type,
			Title:        rule.Title,
			Link:         rule.Link,
			Path:         rule.Path,
			Icon:         rule.Icon,
			MenuType:     rule.MenuType,
			Remark:       rule.Remark,
			Active:       rule.Active,
			Sequence:     rule.Sequence,
			CreatedAt:    string(rule.CreatedAt),
			UpdatedAt:    string(rule.UpdatedAt),
		})
	}

	return _ruleChildren("", pdata)
}

func _ruleChildren(pid string, pdata map[string][]RuleTreeNode) (tree []RuleTree) {
	for _, rule := range pdata[pid] {
		var _tree RuleTree
		_tree.RuleTreeNode = rule
		_tree.Children = _ruleChildren(rule.Nanoid, pdata)

		tree = append(tree, _tree)
	}
	return tree
}

func (l *_rule) AdminRuleTree() (result any, err error) {
	var ruleData []*entity.AdminRule

	ruleData, err = entity.Rule.GetRuleAll()

	if len(ruleData) <= 0 || err != nil {
		err = errors.New("权限节点无数据")
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
