package models

type Pagination struct {
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	Sort      string      `json:"sort"`
	TotalRows int64       `json:"totalRows"`
	Results   interface{} `json:"results"`
}
