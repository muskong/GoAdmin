package entity

import (
	"fmt"

	"github.com/muskong/GoPkg/sqlx"
	"github.com/muskong/GoPkg/zaplog"
)

const (
	AdminRoleRuleTable = `admin_role_rules`
)

type (
	AdminRoleRuleObject struct {
		RoleId int64 `json:"roleId,omitempty" db:"role_id"`
		RuleId int64 `json:"ruleId,omitempty" db:"rule_id"`
	}
)

func GetRuleAllByRoleId(roleId int64) (roleRules []AdminRoleRuleObject) {
	var sqlString string = "SELECT * FROM %s WHERE role_id = ?"
	sql := fmt.Sprintf(sqlString, AdminRoleRuleTable)

	err := sqlx.Select(&roleRules, sql, roleId)
	if err != nil {
		zaplog.Sugar.Errorf("查询出错，%v", err)
	}
	return
}

func GetRoleAllByRuleId(ruleId int64) (roleRules []AdminRoleRuleObject) {
	var sqlString string = "SELECT * FROM %s WHERE rule_id = ?"
	sql := fmt.Sprintf(sqlString, AdminRoleRuleTable)

	err := sqlx.Select(&roleRules, sql, ruleId)
	if err != nil {
		zaplog.Sugar.Errorf("查询出错，%v", err)
	}
	return
}

func InsertAdminRoleRule(roleRules AdminRoleRuleObject) int64 {
	// role.Roles = ""

	return sqlx.NamedUpdate(fmt.Sprintf(`INSERT INTO %s
		SET role_id=:role_id,
			rule_id=:rule_id`, AdminRoleRuleTable), roleRules)
}

func DeleteAdminRoleRule(roleRules AdminRoleRuleObject) int64 {
	return sqlx.NamedUpdate(fmt.Sprintf(`DELETE FROM %s WHERE role_id=:role_id AND rule_id=:rule_id`, AdminRoleRuleTable), roleRules)
}
