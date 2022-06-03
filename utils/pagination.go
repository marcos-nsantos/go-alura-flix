package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/models"
	"strconv"
)

func GeneratePagination(c *gin.Context, model interface{}) *models.Pagination {
	limit := 5
	page := 1
	sort := "created_at asc"

	query := c.Request.URL.Query()
	if query.Get("limit") != "" {
		limit, _ = strconv.Atoi(query.Get("limit"))
	}
	if query.Get("page") != "" {
		page, _ = strconv.Atoi(query.Get("page"))
	}
	if query.Get("sort") != "" {
		sort = query.Get("sort")
	}

	return &models.Pagination{
		Page:    page,
		Limit:   limit,
		Sort:    sort,
		Results: model,
	}
}
