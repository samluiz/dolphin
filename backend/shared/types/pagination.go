package types

import (
	"errors"
	"math"
)

type PaginatedResult struct {
	Pagination PaginationOutput `json:"pagination"`
	Data       interface{}      `json:"data"`
}

type PaginationOutput struct {
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	TotalPages int    `json:"total_pages"`
	TotalItems int    `json:"total_items"`
	NextPage   int    `json:"next_page"`
	PrevPage   int    `json:"prev_page"`
	OrderBy    string `json:"order_by"`
	SortBy     string `json:"sort_by"`
}

type Pagination struct {
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	OrderBy string `json:"order_by"`
	SortBy  string `json:"sort_by"`
}

type DatabasePagination struct {
	Offset     int
	Limit      int
	TotalPages int
	OrderBy    string
	SortBy     string
	NextPage   int
	PrevPage   int
}

// Validate checks the pagination values and returns the total pages and an error if there is one
func (p *Pagination) validate(totalItems int) (int, error) {
	if p.Size <= 0 {
		return 0, ErrSizeOutOfRange
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(p.Size)))
	if totalPages == 0 {
		totalPages = 1
	}
	if p.Page > totalPages {
		return 0, ErrPageOutOfRange
	}
	return totalPages, nil
}

// GetValues computes pagination values and returns them in DatabasePagination
func (p *Pagination) GetValues(totalItems int) (DatabasePagination, error) {
	if p.Page <= 0 {
		p.Page = 1
	}

	totalPages, err := p.validate(totalItems)
	if err != nil {
		return DatabasePagination{}, err
	}

	offset := (p.Page - 1) * p.Size

	orderBy := p.OrderBy
	if orderBy == "" {
		orderBy = "created_at"
	}

	sortBy := p.SortBy
	if sortBy == "" {
		sortBy = "DESC"
	}

	nextPage := 0
	if p.Page < totalPages {
		nextPage = p.Page + 1
	}

	prevPage := 0
	if p.Page > 1 {
		prevPage = p.Page - 1
	}

	return DatabasePagination{
		Offset:     offset,
		Limit:      p.Size,
		TotalPages: totalPages,
		OrderBy:    orderBy,
		SortBy:     sortBy,
		NextPage:   nextPage,
		PrevPage:   prevPage,
	}, nil
}

var (
	ErrPageOutOfRange = errors.New("page out of range")
	ErrSizeOutOfRange = errors.New("size out of range")
)