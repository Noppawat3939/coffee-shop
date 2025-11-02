package helpers

import "gorm.io/gorm"

type Pagination struct {
	Limit int
	Page  int
}

const (
	DefaultPage  = 1
	DefaultLimit = 10
)

func (p *Pagination) getPage() int {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}

	return p.Page
}

func (p *Pagination) getLimit() int {
	if p.Limit > 50 || p.Limit <= 0 {
		p.Limit = DefaultLimit
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
