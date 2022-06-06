package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/models"
	"strconv"
)

func GeneratePagination[T models.Models](c *gin.Context) models.Pagination[T] {
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

	return models.Pagination[T]{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}
}
