package util

type Page struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (p Page) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

func (p Page) GetPage() int {
	return p.Page
}
