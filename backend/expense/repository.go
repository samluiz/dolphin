package expense

import (
	t "fintrack/backend/shared/types"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindAll() ([]t.ExpenseOutput, error)
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
		"UPDATE expenses SET description = ?, amount = ? WHERE id = ? RETURNING *",
		e.Description, e.Amount, id)

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

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}
