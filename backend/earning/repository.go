package earning

import (
	t "dolphin/backend/shared/types"
	"fmt"
	"time"

	p "github.com/Saurs-Developers/go-pagination"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindAll() ([]t.EarningOutput, error)
	FindAllByProfileID(profileID int, pagination p.Pagination) (p.GenericPaginatedResult, error)
	FindByID(id int) (t.EarningToUpdate, error)
	Create(e t.EarningInput) (t.Earning, error)
	Update(id int, e t.EarningUpdate) (t.Earning, error)
	Delete(id int) error
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) FindAll() ([]t.EarningOutput, error) {
	var earnings []t.EarningOutput

	err := r.db.Select(&earnings, "SELECT id, description, amount, (SELECT SUM(amount) FROM earnings) AS sub_total FROM earnings")

	if err != nil {
		return nil, err
	}

	return earnings, nil
}

func (r *repository) FindAllByProfileID(profileID int, pagination p.Pagination) (p.GenericPaginatedResult, error) {
	var earnings []t.EarningOutput

	total, err := r.countByProfileID(profileID)

	if err != nil {
		return p.GenericPaginatedResult{}, err
	}

	dbPagination, err := pagination.GetValues(total)

	if err != nil {
		return p.GenericPaginatedResult{}, err
	}

	validOrderBys := map[string]bool{"id": true, "description": true, "amount": true, "created_at": true, "updated_at": true}
	validSortBys := map[string]bool{"ASC": true, "DESC": true, "asc": true, "desc": true}

	if !validOrderBys[dbPagination.OrderBy] {
		return p.GenericPaginatedResult{}, fmt.Errorf("invalid order by value: %s", dbPagination.OrderBy)
	}

	if !validSortBys[dbPagination.SortBy] {
		return p.GenericPaginatedResult{}, fmt.Errorf("invalid sort by value: %s", dbPagination.SortBy)
	}

	query := fmt.Sprintf(`
	SELECT e.id, e.description, e.amount, (SELECT SUM(amount) FROM earnings WHERE profile_id = e.profile_id) AS sub_total, e.created_at, e.updated_at FROM earnings AS e WHERE profile_id = ?
	ORDER BY e.%s %s
	LIMIT ? OFFSET ?`, dbPagination.OrderBy, dbPagination.SortBy)

	err = r.db.Select(&earnings, query, profileID, dbPagination.Limit, dbPagination.Offset)
	if err != nil {
		return p.GenericPaginatedResult{}, err
	}

	result := p.GenericPaginatedResult{
		Pagination: p.PaginationOutput{
			Page:       pagination.Page,
			Size:       pagination.Size,
			TotalPages: dbPagination.TotalPages,
			TotalItems: total,
			NextPage:  dbPagination.NextPage,
			PrevPage:  dbPagination.PrevPage,
			OrderBy:   dbPagination.OrderBy,
			SortBy:    dbPagination.SortBy,
	},
		Data: earnings,
	}

	return result, nil
}
		
func (r *repository) FindByID(id int) (t.EarningToUpdate, error) {
	var earning t.EarningToUpdate

	err := r.db.Get(&earning, "SELECT description, amount FROM earnings WHERE id = ?", id)

	if err != nil {
		return t.EarningToUpdate{}, err
	}

	return earning, nil
}

func (r *repository) Create(e t.EarningInput) (t.Earning, error) {
	var earning t.Earning

	err := r.db.Get(&earning,
		"INSERT INTO earnings (description, amount, profile_id) VALUES (?, ?, ?) RETURNING *",
		e.Description, e.Amount, e.ProfileID)

	if err != nil {
		return t.Earning{}, err
	}

	return earning, nil
}

func (r *repository) Update(id int, e t.EarningUpdate) (t.Earning, error) {
	var earning t.Earning

	err := r.db.Get(&earning,
		"UPDATE earnings SET description = ?, amount = ?, updated_at = ? WHERE id = ? RETURNING *",
		e.Description, e.Amount, time.Now().Format("2006-01-02 15:04:05"), id)

	if err != nil {
		return t.Earning{}, err
	}

	return earning, nil
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM earnings WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) countByProfileID(profileID int) (int, error) {
	var total int

	err := r.db.Get(&total, "SELECT COUNT(*) FROM earnings WHERE profile_id = ?", profileID)

	if err != nil {
		return 0, err
	}

	return total, nil
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}
