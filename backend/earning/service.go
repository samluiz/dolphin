package earning

import t "dolphin/backend/shared/types"

type Service interface {
	FindAll() ([]t.EarningOutput, error)
	FindAllByProfileID(profileID int) ([]t.EarningOutput, error)
	FindByID(id int) (t.EarningToUpdate, error)
	Create(e t.EarningInput) (t.Earning, error)
	Update(id int, e t.EarningUpdate) (t.Earning, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func (s *service) FindAll() ([]t.EarningOutput, error) {
	return s.r.FindAll()
}

func (s *service) FindAllByProfileID(profileID int) ([]t.EarningOutput, error) {
	return s.r.FindAllByProfileID(profileID)
}

func (s *service) FindByID(id int) (t.EarningToUpdate, error) {
	return s.r.FindByID(id)
}

func (s *service) Create(e t.EarningInput) (t.Earning, error) {
	return s.r.Create(e)
}

func (s *service) Update(id int, e t.EarningUpdate) (t.Earning, error) {
	return s.r.Update(id, e)
}

func (s *service) Delete(id int) error {
	return s.r.Delete(id)
}

func NewService(r Repository) Service {
	return &service{r}
}
