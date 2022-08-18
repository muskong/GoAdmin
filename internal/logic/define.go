package logic

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
)

type (
	Logic struct {
		Ctx *gin.Context
	}
	Page struct {
		Page  int `json:"Page"`
		Limit int `json:"Limit"`
	}
	Result struct {
		Data       interface{}
		Pagination struct {
			Total int64
			Page
		}
	}
	SelectInterface struct {
		Value   string `json:"Value"`
		Name    string `json:"Name"`
		Checked bool   `json:"Checked"`
	}

	RoleRuleObject struct {
		RoleId  string   `json:"RoleId"`
		RuleIds []string `json:"RuleIds"`
	}
	RuleTreeNode struct {
		Id           int    `json:"Id"`
		Nanoid       string `json:"Nanoid"`
		ParentNanoid string `json:"ParentNanoid"`
		Type         string `json:"Type"`
		Title        string `json:"Title"`
		Link         string `json:"Name"`
		Path         string `json:"Path"`
		Icon         string `json:"Icon"`
		MenuType     string `json:"MenuType"`
		Remark       string `json:"Remark"`
		Active       string `json:"Active"`
		Sequence     string `json:"Sequence"`
		CreatedAt    string `json:"CreatedAt"`
		UpdatedAt    string `json:"UpdatedAt"`
	}

	RuleTree struct {
		RuleTreeNode
		Children []RuleTree `json:"Children,omitempty"`
	}

	RoleTreeNode struct {
		Id           int    `json:"Id"`
		Nanoid       string `json:"Nanoid"`
		ParentNanoid string `json:"ParentNanoid"`
		Name         string `json:"Name"`
		Description  string `json:"Description"`
		State        string `json:"State"`
		CreatedAt    string `json:"CreatedAt,omitempty" db:"created_at"`
		UpdatedAt    string `json:"UpdatedAt,omitempty" db:"updated_at"`
	}

	RoleTree struct {
		RoleTreeNode
		Children []RoleTree `json:"Children,omitempty"`
	}

	AntTreeNode struct {
		Title string `json:"title"`
		Value string `json:"value"`
	}
	AntTreeSelect struct {
		AntTreeNode
		Children []AntTreeSelect `json:"children,omitempty"`
	}
)

func (l *Logic) Context(ctx *gin.Context) {
	l.Ctx = ctx
}

func (l *Logic) Log(remark string, data any) {
	param, _ := json.Marshal(data)
	userId, _ := l.Ctx.Get("userId")
	userName, _ := l.Ctx.Get("userName")
	entity.AdminLog.InsertAdminLog(&entity.AdminLogEntity{
		AdminId:   userId.(string),
		AdminName: userName.(string),
		Ip:        l.Ctx.ClientIP(),
		Url:       l.Ctx.Request.RequestURI,
		Method:    l.Ctx.Request.Method,
		Type:      l.Ctx.Request.Header.Get("Content-Type"),
		Useragent: l.Ctx.Request.Header.Get("User-Agent"),
		Param:     string(param),
		Remark:    remark,
	})
}

func (p *Page) getOffset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Limit * (p.Page - 1)
}

func AntdTree(parentNanoid string, pdata map[string][]AntTreeNode) (tree []AntTreeSelect) {
	for _, role := range pdata[parentNanoid] {
		var _tree AntTreeSelect
		_tree.AntTreeNode = role
		_tree.Children = AntdTree(role.Value, pdata)

		tree = append(tree, _tree)
	}
	return tree
}
