package entity

import (
	"database/sql"
	"fmt"

	"github.com/muskong/GoPkg/sqlx"
	"github.com/muskong/GoPkg/zaplog"
)

const (
	AdminRoleTable = `admin_roles`
)

type (
	AdminRoleObject struct {
		AdminRoleId          int64          `json:"adminRoleId,omitempty" db:"admin_role_id"`
		AdminRoleName        string         `json:"adminRoleName,omitempty" db:"admin_role_name"`
		AdminRoleDescription string         `json:"adminRoleDescription,omitempty" db:"admin_role_description"`
		CreatedAt            string         `json:"createdAt,omitempty" db:"created_at"`
		UpdatedAt            sql.NullString `json:"updatedAt,omitempty" db:"updated_at"`
		DeletedAt            sql.NullString `json:"deletedAt,omitempty" db:"deleted_at"`
	}
)

func GetRole(roleId int64) (role AdminRoleObject) {
	field := "*"
	where := "admin_role_id=?"

	sql := sqlx.SqlDeleted(field, AdminRoleTable, where, "admin_role_id", "ASC", 1, 0)

	err := sqlx.Get(&role, sql, roleId)

	if err != nil {
		zaplog.Sugar.Errorf("没有找到，%v", err)
	}
	return
}

func GetRoleAll() (roles []AdminRoleObject) {
	var sqlString string = "SELECT * FROM %s WHERE deleted_at IS NULL"
	sql := fmt.Sprintf(sqlString, AdminRoleTable)

	err := sqlx.Select(&roles, sql)
	if err != nil {
		zaplog.Sugar.Errorf("查询出错，%v", err)
	}
	return
}

// 按选项查询集合
// offset 跳过
// limit 读取数量
// field  排序字段
// sort  排序   DESC 倒叙 ， ASC 正序
func GetRoles(offset, limit int64, sortField, sort string) (roles []AdminRoleObject) {
	field := "*"
	where := "admin_role_id>0"
	sql := sqlx.SqlDeleted(field, AdminRoleTable, where, sortField, sort, limit, offset)

	err := sqlx.Select(&roles, sql)
	if err != nil {
		zaplog.Sugar.Errorf("查询出错，%v", err)
	}
	return
}

func GetRoleCount() int64 {
	var total int64 = 0

	field := "count(1) as total"
	where := "admin_role_id>0"

	sql := sqlx.SqlDeleted(field, AdminRoleTable, where, "admin_role_id", "ASC", 1, 0)

	err := sqlx.Get(&total, sql)

	if err != nil {
		zaplog.Sugar.Errorf("没有找到，%v", err)
	}

	return total
}

func InsertAdminRole(role AdminRoleObject) int64 {
	// role.Roles = ""

	return sqlx.NamedInsert(fmt.Sprintf(`INSERT INTO %s
		SET admin_role_name=:admin_role_name,
			admin_role_description=:admin_role_description`, AdminRoleTable), role)
}
