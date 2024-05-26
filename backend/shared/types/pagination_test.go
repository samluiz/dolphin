package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaginationValidate(t *testing.T) {
	tests := []struct {
		name       string
		pagination Pagination
		totalItems int
		expected   int
		expectErr  error
	}{
		{"Valid pagination", Pagination{Page: 1, Size: 10}, 100, 10, nil},
		{"Zero totalItems", Pagination{Page: 1, Size: 10}, 0, 1, nil},
		{"Page out of range", Pagination{Page: 5, Size: 10}, 20, 0, ErrPageOutOfRange},
		{"Negative page", Pagination{Page: -1, Size: 10}, 100, 10, nil},
		{"Zero size", Pagination{Page: 1, Size: 0}, 100, 0, ErrSizeOutOfRange},
		{"Negative size", Pagination{Page: 1, Size: -10}, 100, 0, ErrSizeOutOfRange},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalPages, err := tt.pagination.validate(tt.totalItems)
			assert.Equal(t, tt.expectErr, err)
			assert.Equal(t, tt.expected, totalPages)
		})
	}
}

func TestPaginationGetValues(t *testing.T) {
	tests := []struct {
		name       string
		pagination Pagination
		totalItems int
		expected   DatabasePagination
		expectErr  error
	}{
		{
			"Valid pagination",
			Pagination{Page: 1, Size: 10, OrderBy: "description", SortBy: "ASC"},
			100,
			DatabasePagination{Offset: 0, Limit: 10, TotalPages: 10, OrderBy: "description", SortBy: "ASC", NextPage: 2, PrevPage: 0},
			nil,
		},
		{
			"Default order and sort",
			Pagination{Page: 1, Size: 10},
			100,
			DatabasePagination{Offset: 0, Limit: 10, TotalPages: 10, OrderBy: "created_at", SortBy: "DESC", NextPage: 2, PrevPage: 0},
			nil,
		},
		{
			"Zero totalItems",
			Pagination{Page: 1, Size: 10},
			0,
			DatabasePagination{Offset: 0, Limit: 10, TotalPages: 1, OrderBy: "created_at", SortBy: "DESC", NextPage: 0, PrevPage: 0},
			nil,
		},
		{
			"Page out of range",
			Pagination{Page: 5, Size: 10},
			20,
			DatabasePagination{},
			ErrPageOutOfRange,
		},
		{
			"Zero page",
			Pagination{Page: 0, Size: 10},
			100,
			DatabasePagination{Offset: 0, Limit: 10, TotalPages: 10, OrderBy: "created_at", SortBy: "DESC", NextPage: 2, PrevPage: 0},
			nil,
		},
		{
			"Negative page",
			Pagination{Page: -1, Size: 10},
			100,
			DatabasePagination{Offset: 0, Limit: 10, TotalPages: 10, OrderBy: "created_at", SortBy: "DESC", NextPage: 2, PrevPage: 0},
			nil,
		},
		{
			"Zero size",
			Pagination{Page: 1, Size: 0},
			100,
			DatabasePagination{},
			ErrSizeOutOfRange,
		},
		{
			"Negative size",
			Pagination{Page: 1, Size: -10},
			100,
			DatabasePagination{},
			ErrSizeOutOfRange,
		},
		{
			"Next and prev pages",
			Pagination{Page: 2, Size: 10},
			100,
			DatabasePagination{Offset: 10, Limit: 10, TotalPages: 10, OrderBy: "created_at", SortBy: "DESC", NextPage: 3, PrevPage: 1},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.pagination.GetValues(tt.totalItems)
			assert.Equal(t, tt.expectErr, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
