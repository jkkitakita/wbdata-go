package wbdata

type PageParams struct {
	Page    int
	PerPage int
}

type PageSummary struct {
	Page    int	`json:"page"`
	Pages   int	`json:"pages"`
	PerPage int `json:"per_page,string"`
	Total   int	`json:"total"`
}
