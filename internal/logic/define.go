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
		Value   int64  `json:"value"`
		Name    string `json:"name"`
		Checked bool   `json:"checked"`
	}

	RoleRuleObject struct {
		RoleId  int64   `json:"roleId"`
		RuleIds []int64 `json:"ruleIds"`
	}
)

func (p *Page) getOffset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Limit * (p.Page - 1)
}
