package types

import (
	"errors"
	"fmt"
	"math"
)

type PaginatedResult struct {
	Pagination PaginationOutput `json:"pagination"`
	Data       interface{} `json:"data"`
}

type PaginationOutput struct {
	Page       int  `json:"page"`
	Size       int	`json:"size"`
	TotalPages int	`json:"total_pages"`
	TotalItems int `json:"total_items"`
	NextPage   int `json:"next_page"`
	PrevPage   int `json:"prev_page"`
	OrderBy    string `json:"order_by"`
	SortBy     string `json:"sort_by"`
}

type Pagination struct {
	Page    int `json:"page"`
	Size    int `json:"size"`
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

// this function is used to validate the pagination values and return the total pages with an error if there is one
func (p *Pagination) validate(totalItems int) (int, error) {

	if (totalItems == 0) {
		return 0, nil
	}

	totalPages := int(math.Round(float64(totalItems) / float64(p.Size)))

	fmt.Println("totalItems: ", totalItems)
	fmt.Println("totalPages: ", totalPages)

	if totalPages == 0 {
		return 1, nil
	}
	if p.Page < 0 {
		return 1, nil
	}
	if p.Size <= 0 {
		return totalPages, ErrSizeOutOfRange
	}
	if totalPages < p.Page {
		return totalPages, ErrPageOutOfRange
	}
	return totalPages, nil
}

// this function is used to get the values for the pagination and return the
// offset, limit, totalPages, orderBy, sortBy and an error if there is one
func (p *Pagination) GetValues(totalItems int) (DatabasePagination, error) {
	var offset int
	var limit int
	var orderBy string
	var sortBy string
	var nextPage int
	var prevPage int

	totalPages, err := p.validate(totalItems)

	if err != nil {
		return DatabasePagination{}, err
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

	if p.Page < totalPages {
		nextPage = p.Page + 1
	} else {
		nextPage = 0
	}

	if p.Page > 1 {
		prevPage = p.Page - 1
	} else {
		prevPage = 0
	}

	return DatabasePagination{
		Offset:     offset,
		Limit:      limit,
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