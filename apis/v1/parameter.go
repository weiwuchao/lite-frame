package v1

type UserSelector struct {
	Name string
}

type Page struct {
	PageNumber int    `json:"page" query:"page" form:"page"`
	PageSize   int    `json:"size" query:"size" form:"size"`
	OrderBy    string `json:"orderBy" query:"orderBy" form:"orderBy"`
	Desc       bool   `json:"desc" query:"desc" form:"desc"`
}

func (p Page) PackOrderSql() string {
	if p.Desc {
		return p.OrderBy + " desc"
	}
	return p.OrderBy
}
