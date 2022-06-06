package models

type Models interface {
	Video | Categoria | User
}

type Pagination[T Models] struct {
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Sort      string `json:"sort"`
	TotalRows int64  `json:"totalRows"`
	Results   *[]T   `json:"results"`
}
