package expense

import (
	t "dolphin/backend/shared/types"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindAll() ([]t.ExpenseOutput, error)
	FindAllByProfileID(profileID int, pagination t.Pagination) (t.PaginatedResult, error)
	FindByID(id int) (t.ExpenseToUpdate, error)
	Create(e t.ExpenseInput) (t.Expense, error)
	Update(id int, e t.ExpenseUpdate) (t.Expense, error)
	Delete(id int) error
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) FindAll() ([]t.ExpenseOutput, error) {
	var expenses []t.ExpenseOutput

	err := r.db.Select(&expenses, `
	SELECT 
		e.id,
		e.description,
		e.amount,
		c.description AS category,
		(SELECT SUM(amount) FROM expenses) AS sub_total
	FROM 
		expenses AS e
	JOIN 
		expense_category AS ec ON e.id = ec.expense_id
	JOIN 
		categories AS c ON ec.category_id = c.id`)

	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *repository) FindAllByProfileID(profileID int, pagination t.Pagination) (t.PaginatedResult, error) {
	var expenses []t.ExpenseOutput

	total, err := r.countByProfileID(profileID)

	if err != nil {
		return t.PaginatedResult{}, err
	}

	dbPagination, err := pagination.GetValues(total)

	if err != nil {
		return t.PaginatedResult{}, err
	}

	validOrderBys := map[string]bool{"id": true, "description": true, "amount": true, "category": true, "created_at": true, "updated_at": true}
	validSortBys := map[string]bool{"ASC": true, "DESC": true, "asc": true, "desc": true}

	if !validOrderBys[dbPagination.OrderBy] {
		return t.PaginatedResult{}, fmt.Errorf("invalid order by value: %s", dbPagination.OrderBy)
	}

	if !validSortBys[dbPagination.SortBy] {
		return t.PaginatedResult{}, fmt.Errorf("invalid sort by value: %s", dbPagination.SortBy)
	}

	query := fmt.Sprintf(`
	SELECT
		e.id,
		e.description,
		e.amount,
		c.description AS category,
		(SELECT SUM(amount) FROM expenses WHERE profile_id = e.profile_id) AS sub_total,
		e.created_at, e.updated_at
	FROM
		expenses AS e
	JOIN
		expense_category AS ec ON e.id = ec.expense_id
	JOIN
		categories AS c ON ec.category_id = c.id
	WHERE
		e.profile_id = ?
	ORDER BY e.%s %s
	LIMIT ? OFFSET ?`, dbPagination.OrderBy, dbPagination.SortBy)

	err = r.db.Select(&expenses, query, profileID, dbPagination.Limit, dbPagination.Offset)

	if err != nil {
		return t.PaginatedResult{}, err
	}

	result := t.PaginatedResult{
		Pagination: t.PaginationOutput{
			Page:       pagination.Page,
			Size:       pagination.Size,
			TotalPages: dbPagination.TotalPages,
			TotalItems: total,
			NextPage:  dbPagination.NextPage,
			PrevPage:  dbPagination.PrevPage,
			OrderBy:   dbPagination.OrderBy,
			SortBy:    dbPagination.SortBy,
	},
		Data: expenses,
	}

	return result, nil
}

func (r *repository) FindByID(id int) (t.ExpenseToUpdate, error) {
	var expense t.ExpenseToUpdate

	err := r.db.Get(&expense, `
	SELECT 
		e.description,
		e.amount,
		c.description AS category
	FROM 
		expenses AS e
	JOIN 
		expense_category AS ec ON e.id = ec.expense_id
	JOIN 
		categories AS c ON ec.category_id = c.id
	WHERE
		e.id = ?
		`, id)

	if err != nil {
		return t.ExpenseToUpdate{}, err
	}

	return expense, nil
}

func (r *repository) Create(e t.ExpenseInput) (t.Expense, error) {
	var expense t.Expense

	err := r.db.Get(&expense,
		"INSERT INTO expenses (description, amount, profile_id) VALUES (?, ?, ?) RETURNING *",
		e.Description, e.Amount, e.ProfileID)

	if err != nil {
		return t.Expense{}, err
	}

	if e.CategoryID != 0 {
		_, err := r.db.Exec("INSERT INTO expense_category (category_id, expense_id) VALUES (?, ?)", e.CategoryID, expense.ID)

		if err != nil {
			return t.Expense{}, err
		}
	}

	return expense, nil
}

func (r *repository) Update(id int, e t.ExpenseUpdate) (t.Expense, error) {
	var expense t.Expense

	err := r.db.Get(&expense,
		"UPDATE expenses SET description = ?, amount = ?, updated_at = ? WHERE id = ? RETURNING *",
		e.Description, e.Amount, time.Now().Format("2006-01-02 15:04:05"), id)

	if err != nil {
		return t.Expense{}, err
	}

	if e.CategoryID != 0 {
		_, err := r.db.Exec("UPDATE expense_category SET category_id = ? WHERE expense_id = ?", e.CategoryID, id)

		if err != nil {
			return t.Expense{}, err
		}
	}

	return expense, nil
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM expenses WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) countByProfileID(profileID int) (int, error) {
	var total int

	err := r.db.Get(&total, "SELECT COUNT(*) FROM expenses WHERE profile_id = ?", profileID)

	if err != nil {
		return 0, err
	}

	return total, nil
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}
