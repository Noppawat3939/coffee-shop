package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit int
	Page  int
}

const (
	DefaultPage  = 1
	DefaultLimit = 50
	MaxLimit     = 50
)

func (p *Pagination) getPage() int {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}

	return p.Page
}

func (p *Pagination) getLimit() int {
	if p.Limit <= 0 {
		p.Limit = DefaultLimit
	}
	if p.Limit > MaxLimit {
		p.Limit = MaxLimit
	}

	return p.Limit
}

func (p *Pagination) getOffset() int {
	return (p.getPage() - 1) * p.getLimit()
}

func (p *Pagination) GetPaginationResult(db *gorm.DB) *gorm.DB {
	offet := p.getOffset()
	pageSize := p.getLimit()

	return db.Offset(offet).Limit(pageSize)
}

func BuildPagination(c *gin.Context) (int, int) {
	page := ToInt(c.DefaultQuery("page", fmt.Sprint(DefaultPage)))
	limit := ToInt(c.DefaultQuery("limit", fmt.Sprint(DefaultLimit)))

	return page, limit
}
