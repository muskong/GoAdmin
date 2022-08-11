package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _role struct{}

var Role = &_role{}

func (*_role) AdminRole(roleId string) (role *entity.AdminRole, err error) {
	role, err = entity.Role.GetRole(roleId)
	if err != nil {
		err = errors.New("获取失败")
		return
	}
	return
}

func (*_role) AdminRoleList(page Page) (err error, result Result) {

	userData, userCount, err := entity.Role.GetRoles(page.getOffset(), page.Limit)
	if len(userData) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	result.Data = userData
	result.Pagination.Page = page
	result.Pagination.Total = userCount

	return
}
func (*_role) AdminRoleCreate(u entity.AdminRole) error {
	err := entity.Role.InsertAdminRole(&u)
	if u.Id <= 0 || err != nil {
		return errors.New("新增失败")
	}
	return err
}
func (*_role) AdminRoleUpdate(u entity.AdminRole) error {
	err := entity.Role.UpdateAdminRole(&u)
	if u.Id <= 0 || err != nil {
		return errors.New("更新失败")
	}
	return err
}
func (*_role) AdminRoleDelete(roleId string) error {
	err := entity.Role.DeleteAdminRole(roleId)
	if err != nil {
		return errors.New("删除失败")
	}
	return err
}

func (*_role) AdminRoleAll() (err error, result []*entity.AdminRole) {

	result, err = entity.Role.GetRoleAll()
	if len(result) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	return
}

func (*_role) AdminRoleRuleList(roleId string) (err error, result []SelectInterface) {

	ruleData, err := entity.Rule.GetRuleAll()
	if len(ruleData) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	roleRuleSelected := map[string]bool{}
	rule, err := entity.Role.GetRole(roleId)
	if rule.Id > 0 && err == nil {
		for _, v := range rule.Rules {
			roleRuleSelected[v] = true
		}
	}

	for _, rule := range ruleData {
		ruleSelect := SelectInterface{
			rule.Nanoid,
			rule.Title,
			false}
		if roleRuleSelected[rule.Nanoid] {
			ruleSelect.Checked = true
		}
		result = append(result, ruleSelect)
	}

	return
}

func (*_role) AdminRoleSaveRule(ou RoleRuleObject) error {
	err := entity.Role.UpdateAdminRoleRules(ou.RoleId, ou.RuleIds)
	if err != nil {
		return errors.New("设置失败")
	}
	return err
}

func (*_role) AdminRoleGroupList() (result []RoleTree, err error) {
	roleData, err := entity.Role.GetRoleAll()
	if len(roleData) <= 0 || err != nil {
		err = errors.New("无角色数据")
		return
	}

	result = _roleTree(roleData, 0)

	return
}

func _roleTree(roles []*entity.AdminRole, pid int) []RoleTree {

	pdata := map[int][]RoleTreeNode{}

	for _, role := range roles {

		pdata[role.Pid] = append(pdata[role.Pid], RoleTreeNode{
			Id:          role.Id,
			Pid:         role.Pid,
			Name:        role.Name,
			Description: role.Description,
			State:       role.State,
			CreatedAt:   string(role.CreatedAt),
			UpdatedAt:   string(role.UpdatedAt),
		})
	}

	return _roleChildren(pid, pdata)
}

func _roleChildren(pid int, pdata map[int][]RoleTreeNode) (tree []RoleTree) {
	for _, role := range pdata[pid] {
		var _tree RoleTree
		_tree.RoleTreeNode = role
		_tree.Children = _roleChildren(role.Id, pdata)

		tree = append(tree, _tree)
	}
	return tree
}

func (*_role) AdminRoleTree() (result []AntTreeSelect, err error) {
	roleData, err := entity.Role.GetRoleAll()
	if len(roleData) <= 0 || err != nil {
		err = errors.New("无角色数据")
		return
	}

	pdata := map[string][]AntTreeNode{}

	for _, item := range roleData {
		pdata[item.ParentNanoid] = append(pdata[item.ParentNanoid], AntTreeNode{
			Value: item.Nanoid,
			Title: item.Name,
		})
	}

	result = AntdTree("", pdata)

	return
}
