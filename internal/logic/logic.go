package logic

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoPkg/zaplog"
	"github.com/spf13/cast"
)

type (
	Api interface {
		Info(dest *[]any) error                                       // 安装
		Send(request map[string]any, respond *map[string]any) error   // 提交数据
		Search(request map[string]any, respond *map[string]any) error // 查询数据
		Notify(request map[string]any, respond *map[string]any) error // 回调处理
	}
	ApiInfo struct {
		Key   string
		Title string
		Value string
	}

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
		Path         string `json:"Path"`
		Icon         string `json:"Icon"`
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

func (l *Logic) Log(title string, data any) {
	param, _ := json.Marshal(data)
	userId, ok := l.Ctx.Get("userId")
	if !ok {
		zaplog.Sugar.Error("adminlog userId")
		userId = ""
	}
	entity.AdminLogEntity.InsertAdminLog(&entity.AdminLogs{
		AdminId:   cast.ToString(userId),
		Ip:        l.Ctx.ClientIP(),
		Url:       l.Ctx.Request.RequestURI,
		Method:    l.Ctx.Request.Method,
		Type:      l.Ctx.Request.Header.Get("Content-Type"),
		Useragent: l.Ctx.Request.Header.Get("User-Agent"),
		Param:     string(param),
		Title:     title,
	})
}

func (p *Page) getLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return p.Limit
}
func (p *Page) getOffset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.getLimit() * (p.Page - 1)
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
