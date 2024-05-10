package category

import (
	t "dolphin/backend/shared/types"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindAll() ([]t.Category, error)
	FindByID(id int) (t.Category, error)
	FindByDescription(description string) (t.Category, error)
	Create(e t.CategoryInput) (t.Category, error)
	Update(id int, e t.CategoryInput) (t.Category, error)
	Delete(id int) error
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) FindAll() ([]t.Category, error) {
	var categories []t.Category

	err := r.db.Select(&categories, "SELECT * FROM categories")

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *repository) FindByID(id int) (t.Category, error) {
	var category t.Category

	err := r.db.Get(&category, "SELECT * FROM categories WHERE id = ?", id)

	if err != nil {
		return t.Category{}, err
	}

	return category, nil
}

func (r *repository) FindByDescription(description string) (t.Category, error) {
	var category t.Category

	err := r.db.Get(&category, "SELECT * FROM categories WHERE description = ?", description)

	if err != nil {
		return t.Category{}, err
	}

	return category, nil
}

func (r *repository) Create(e t.CategoryInput) (t.Category, error) {
	var category t.Category

	err := r.db.Get(&category,
		"INSERT INTO categories (description) VALUES (?) RETURNING *",
		e.Description)

	if err != nil {
		return t.Category{}, err
	}

	return category, nil
}

func (r *repository) Update(id int, e t.CategoryInput) (t.Category, error) {
	var category t.Category

	err := r.db.Get(&category,
		"UPDATE categories SET description = ? WHERE id = ? RETURNING *",
		e.Description, id)

	if err != nil {
		return t.Category{}, err
	}

	return category, nil
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}
