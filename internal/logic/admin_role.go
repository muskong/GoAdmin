package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
)

type _role struct {
	Logic
}

var Role = &_role{}

func (l *_role) AdminRole(roleId string) (role *entity.AdminRole, err error) {
	role, err = entity.Role.GetRole(roleId)
	if err != nil {
		err = errors.New("获取失败")
		return
	}
	return
}

func (l *_role) AdminRoleList(page Page) (result Result, err error) {

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
func (l *_role) AdminRoleCreate(u entity.AdminRole) error {
	err := entity.Role.InsertAdminRole(&u)
	if u.Id <= 0 || err != nil {
		return errors.New("新增角色失败")
	}

	l.Log("新增角色", u)
	return err
}
func (l *_role) AdminRoleUpdate(u entity.AdminRole) error {
	err := entity.Role.UpdateAdminRole(&u)
	if u.Id <= 0 || err != nil {
		return errors.New("更新角色失败")
	}
	l.Log("更新角色", u)
	return err
}
func (l *_role) AdminRoleDelete(roleId string) error {
	err := entity.Role.DeleteAdminRole(roleId)
	if err != nil {
		return errors.New("删除角色失败")
	}
	l.Log("删除角色", roleId)
	return err
}

func (l *_role) AdminRoleAll() (result []*entity.AdminRole, err error) {

	result, err = entity.Role.GetRoleAll()
	if len(result) <= 0 || err != nil {
		err = errors.New("无数据")
		return
	}

	return
}

func (l *_role) AdminRoleRuleList(roleId string) (result []SelectInterface, err error) {

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

func (l *_role) AdminRoleSaveRule(ou RoleRuleObject) error {
	err := entity.Role.UpdateAdminRoleRules(ou.RoleId, ou.RuleIds)
	if err != nil {
		return errors.New("设置失败")
	}
	l.Log("更新角色权限", ou)
	return err
}

func (l *_role) AdminRoleGroupList() (result []RoleTree, err error) {
	roleData, err := entity.Role.GetRoleAll()
	if len(roleData) <= 0 || err != nil {
		err = errors.New("无角色数据")
		return
	}

	pdata := map[string][]RoleTreeNode{}

	for _, role := range roleData {

		pdata[role.ParentNanoid] = append(pdata[role.ParentNanoid], RoleTreeNode{
			Id:           role.Id,
			Nanoid:       role.Nanoid,
			ParentNanoid: role.ParentNanoid,
			Name:         role.Name,
			Description:  role.Description,
			State:        role.State,
			CreatedAt:    string(role.CreatedAt),
			UpdatedAt:    string(role.UpdatedAt),
		})
	}

	result = _roleChildren("", pdata)
	return
}

func _roleChildren(pid string, pdata map[string][]RoleTreeNode) (tree []RoleTree) {
	for _, role := range pdata[pid] {
		var _tree RoleTree
		_tree.RoleTreeNode = role
		_tree.Children = _roleChildren(role.Nanoid, pdata)

		tree = append(tree, _tree)
	}
	return tree
}

func (l *_role) AdminRoleTree() (result []AntTreeSelect, err error) {
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
