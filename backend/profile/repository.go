package profile

import (
	t "dolphin/backend/shared/types"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindAll() ([]t.Profile, error)
	FindByID(id int) (t.Profile, error)
	Create(e t.ProfileInput) (t.Profile, error)
	Update(id int, e t.ProfileInput) (t.Profile, error)
	Delete(id int) error
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) FindAll() ([]t.Profile, error) {
	var profiles []t.Profile

	err := r.db.Select(&profiles, "SELECT * FROM profiles")

	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r *repository) FindByID(id int) (t.Profile, error) {
	var profile t.Profile

	err := r.db.Get(&profile, "SELECT * FROM profiles WHERE id = ?", id)

	if err != nil {
		return t.Profile{}, err
	}

	return profile, nil
}

func (r *repository) Create(e t.ProfileInput) (t.Profile, error) {
	var profile t.Profile

	if e.IsDefault {
		_, err := r.db.Exec("UPDATE profiles SET is_default = ? WHERE is_default = ?", false, true)

		if err != nil {
			return t.Profile{}, err
		}
	}

	err := r.db.Get(&profile,
		"INSERT INTO profiles (description, is_default) VALUES (?, ?) RETURNING *",
		e.Description, e.IsDefault)

	if err != nil {
		return t.Profile{}, err
	}

	return profile, nil
}

func (r *repository) Update(id int, e t.ProfileInput) (t.Profile, error) {
	var profile t.Profile

	profiles, err := r.FindAll()

	if err != nil {
		return t.Profile{}, nil 
	}

	if e.IsDefault {
		for _, p := range profiles {
			if p.ID != id && p.IsDefault {
				_, err := r.db.Exec("UPDATE profiles SET is_default = ? WHERE id = ?", false, p.ID)

				if err != nil {
					return t.Profile{}, err
				}
			}
		}
	}

	err = r.db.Get(&profile,
		"UPDATE profiles SET description = ?, is_default = ?,  updated_at = ? WHERE id = ? RETURNING *",
		e.Description, e.IsDefault, time.Now().Format("2006-01-02 15:04:05"), id)

	if err != nil {
		return t.Profile{}, err
	}

	return profile, nil
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM profiles WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}
