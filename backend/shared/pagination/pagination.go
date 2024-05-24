package pagination

import (
	"errors"
	"math"
)

type Pagination struct {
	Page    int
	Size    int
	OrderBy string
	SortBy  string
}

// this function is used to validate the pagination values and return the total pages with an error if there is one
func (p *Pagination) validate(totalItems int) (int, error) {
	totalPages := int(math.Round(float64(totalItems) / float64(p.Size)))
	if p.Page < 0 {
		return totalPages, ErrPageOutOfRange
	}
	if p.Size < 0 {
		return totalPages, ErrSizeOutOfRange
	}
	if totalPages < p.Page {
		return totalPages, ErrPageOutOfRange
	}
	return totalPages, nil
}

// this function is used to get the values for the pagination and return the 
// offset, limit, totalPages, orderBy, sortBy and an error if there is one
func (p *Pagination) GetValues(totalItems int) (int, int, int, string, string, error) {
	var offset int
	var limit int
	var orderBy string
	var sortBy string

	totalPages, err := p.validate(totalItems)

	if err != nil {
		return 0, 0, totalPages, "", "", err
	}

	if p.Page > 0 {
		offset = (p.Page - 1) * p.Size
	} else {
		offset = 0
	}

	if p.Size > 0 {
		limit = p.Size
	} else {
		limit = 10
	}

	if p.OrderBy == "" {
		orderBy = "id"
	} else {
		orderBy = p.OrderBy
	}
	if p.SortBy == "" {
		sortBy = "ASC"
	} else {
		sortBy = p.SortBy
	}

	return offset, limit, totalPages, orderBy, sortBy, nil
}

var (
	ErrPageOutOfRange = errors.New("page out of range")
	ErrSizeOutOfRange = errors.New("size out of range")
)