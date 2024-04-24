package category

import t "fintrack/backend/shared/types"

type Service interface {
	FindAll() ([]t.Category, error)
	FindByID(id int) (t.Category, error)
	FindByDescription(description string) (t.Category, error)
	Create(e t.CategoryInput) (t.Category, error)
	Update(id int, e t.CategoryInput) (t.Category, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func (s *service) FindAll() ([]t.Category, error) {
	return s.r.FindAll()
}

func (s *service) FindByID(id int) (t.Category, error) {
	return s.r.FindByID(id)
}

func (s *service) FindByDescription(description string) (t.Category, error) {
	return s.r.FindByDescription(description)
}

func (s *service) Create(e t.CategoryInput) (t.Category, error) {
	return s.r.Create(e)
}

func (s *service) Update(id int, e t.CategoryInput) (t.Category, error) {
	return s.r.Update(id, e)
}

func (s *service) Delete(id int) error {
	return s.r.Delete(id)
}

func NewService(r Repository) Service {
	return &service{r}
}
