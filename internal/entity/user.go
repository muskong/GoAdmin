package entity

import (
	"database/sql"
	"fmt"

	"github.com/muskong/GoPkg/sqlx"
	"github.com/muskong/GoPkg/zaplog"
)

const (
	AdminUserTable = `admin_users`
)

type (
	AdminUserObject struct {
		/** ID **/
		AdminUserId int64 `json:"adminUserId,omitempty" db:"admin_user_id"`
		/** 用户名 **/
		AdminUserName string `json:"adminUserName,omitempty" db:"admin_user_name"`
		/** 密码 **/
		AdminUserPassword string `json:"adminUserPassword,omitempty" db:"admin_user_password"`
		/** 角色 **/
		AdminUserRoles sqlx.JsonString `json:"adminUserRoles,omitempty" db:"admin_user_roles"`

		/** 发布时间 **/
		CreatedAt string `json:"createdAt,omitempty" db:"created_at"`
		/** 更新时间 **/
		UpdatedAt sql.NullString `json:"updatedAt,omitempty" db:"updated_at"`
		DeletedAt sql.NullString `json:"deletedAt,omitempty" db:"deleted_at"`
	}
)

func GetAdminName(name string) (admin AdminUserObject) {
	field := "admin_user_id, admin_user_name, admin_user_password"
	where := "admin_user_name=?"

	sql := sqlx.SqlDeleted(field, AdminUserTable, where, "admin_user_id", "ASC", 1, 0)

	err := sqlx.Get(&admin, sql, name)

	if err != nil {
		zaplog.Sugar.Info(sql)
		zaplog.Sugar.Errorf("没有找到[%s]，%v", name, err)
	}
	return
}

func GetUser(adminId int64) (admin AdminUserObject) {
	field := "*"
	where := "admin_user_id=?"

	sql := sqlx.SqlDeleted(field, AdminUserTable, where, "admin_user_id", "ASC", 1, 0)

	err := sqlx.Get(&admin, sql, adminId)

	if err != nil {
		zaplog.Sugar.Errorf("没有找到，%v", err)
	}
	return
}

// 按选项查询集合
// offset 跳过
// limit 读取数量
// field  排序字段
// sort  排序   DESC 倒叙 ， ASC 正序
func GetUsers(offset, limit int64, sortField, sort string) (users []AdminUserObject) {
	field := "*"
	where := "admin_user_id>0"
	sql := sqlx.SqlDeleted(field, AdminUserTable, where, sortField, sort, limit, offset)
	zaplog.Sugar.Info(sql)
	err := sqlx.Select(&users, sql)
	if err != nil {
		zaplog.Sugar.Errorf("查询出错，%v", err)
	}
	return
}

func GetUserCount() int64 {
	var total int64 = 0

	field := "count(1) as total"
	where := "admin_user_id>0"

	sql := sqlx.SqlDeleted(field, AdminUserTable, where, "admin_user_id", "ASC", 1, 0)

	err := sqlx.Get(&total, sql)

	if err != nil {
		zaplog.Sugar.Errorf("没有找到，%v", err)
	}

	return total
}

func InsertAdminUser(user AdminUserObject) int64 {
	// user.Roles = ""

	return sqlx.NamedInsert(fmt.Sprintf(`INSERT INTO %s
		SET admin_user_name=:admin_user_name,
			admin_user_password=:admin_user_password,
			admin_user_roles=:admin_user_roles`, AdminUserTable), user)
}
