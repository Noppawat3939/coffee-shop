package pagination

import (
	"strconv"

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

func NewFromQuery(c *gin.Context) *Pagination {
	page, _ := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(DefaultPage)))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(DefaultLimit)))

	p := &Pagination{
		Page:  page,
		Limit: limit,
	}

	p.normalize()

	return p
}

func (p *Pagination) normalize() {

	if p.Page <= 0 {
		p.Page = DefaultPage
	}

	if p.Limit <= 0 {
		p.Limit = DefaultLimit
	}

	if p.Limit > MaxLimit {
		p.Limit = MaxLimit
	}
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

func (p *Pagination) Apply(db *gorm.DB) *gorm.DB {
	return db.Offset(p.Offset()).Limit(p.Limit)
}
