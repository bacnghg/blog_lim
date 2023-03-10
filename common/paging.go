package common

type Paging struct {
	Page      int `json:"page" form:"page"`
	Limit     int `json:"limit" form:"limit"`
	Total     int64 `json:"total" form:"_"`
	TotalPage float64 `json:"lastPage" form:"_"`
}

func (p *Paging) Validate() error {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
	return nil
}
