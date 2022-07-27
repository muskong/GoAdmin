package logic

type (
	Page struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}
	Result struct {
		Data       interface{}
		Pagination struct {
			Total int64
			Page
		}
	}
	SelectInterface struct {
		Value   int    `json:"value"`
		Name    string `json:"name"`
		Checked bool   `json:"checked"`
	}

	RoleRuleObject struct {
		RoleId  int   `json:"roleId"`
		RuleIds []int `json:"ruleIds"`
	}
	TreeNode struct {
		Id        int    `json:"id"`
		Pid       int    `json:"pid"`
		Type      string `json:"type"`
		Title     string `json:"title"`
		Link      string `json:"name"`
		Path      string `json:"path"`
		Icon      string `json:"icon"`
		MenuType  string `json:"menuType"`
		Url       string `json:"url"`
		Component string `json:"component"`
		Extend    string `json:"extend"`
		KeepAlive string `json:"keepAlive,omitempty"`
	}

	Tree struct {
		TreeNode
		Children []Tree `json:"children,omitempty"`
	}
)

func (p *Page) getOffset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Limit * (p.Page - 1)
}
