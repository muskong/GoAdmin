package logic

type (
	Page struct {
		Page  int64 `json:"page"`
		Limit int64 `json:"limit"`
	}
	Result struct {
		Data       interface{}
		Pagination struct {
			Total int64
			Page  int64
			Limit int64
		}
	}
	SelectInterface struct {
		Value   int64  `json:"value"`
		Name    string `json:"name"`
		Checked bool   `json:"checked"`
	}

	RoleRuleObject struct {
		RoleId  int64 `json:"roleId"`
		RuleId  int64 `json:"ruleId"`
		Checked bool  `json:"checked"`
	}
)

func (p *Page) getOffset() int64 {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Limit * (p.Page - 1)
}
