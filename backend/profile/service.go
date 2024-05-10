package profile

import t "dolphin/backend/shared/types"

type Service interface {
	FindAll() ([]t.Profile, error)
	FindByID(id int) (t.Profile, error)
	Create(e t.ProfileInput) (t.Profile, error)
	Update(id int, e t.ProfileInput) (t.Profile, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func (s *service) FindAll() ([]t.Profile, error) {
	return s.r.FindAll()
}

func (s *service) FindByID(id int) (t.Profile, error) {
	return s.r.FindByID(id)
}

func (s *service) Create(e t.ProfileInput) (t.Profile, error) {
	return s.r.Create(e)
}

func (s *service) Update(id int, e t.ProfileInput) (t.Profile, error) {
	return s.r.Update(id, e)
}

func (s *service) Delete(id int) error {
	return s.r.Delete(id)
}

func NewService(r Repository) Service {
	return &service{r}
}
