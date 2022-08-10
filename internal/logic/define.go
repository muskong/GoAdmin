package logic

type (
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
		Value   int    `json:"Value"`
		Name    string `json:"Name"`
		Checked bool   `json:"Checked"`
	}

	RoleRuleObject struct {
		RoleId  int   `json:"RoleId"`
		RuleIds []int `json:"RuleIds"`
	}
	RuleTreeNode struct {
		Id        int    `json:"Id"`
		Pid       int    `json:"Pid"`
		Type      string `json:"Type"`
		Title     string `json:"Title"`
		Link      string `json:"Name"`
		Path      string `json:"Path"`
		Icon      string `json:"Icon"`
		MenuType  string `json:"MenuType"`
		Url       string `json:"Url"`
		Component string `json:"Component"`
		Extend    string `json:"Extend"`
		KeepAlive string `json:"KeepAlive,omitempty"`
	}

	RuleTree struct {
		RuleTreeNode
		Children []RuleTree `json:"Children,omitempty"`
	}

	RoleTreeNode struct {
		Id          int    `json:"Id"`
		Pid         int    `json:"Pid"`
		Name        string `json:"Name"`
		Description string `json:"Description"`
		State       string `json:"State"`
		CreatedAt   string `json:"CreatedAt,omitempty" db:"created_at"`
		UpdatedAt   string `json:"UpdatedAt,omitempty" db:"updated_at"`
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
