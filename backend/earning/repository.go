package earning

import (
	t "dolphin/backend/shared/types"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindAll() ([]t.EarningOutput, error)
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

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}
