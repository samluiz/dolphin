package expense

import t "fintrack/backend/shared/types"

type Service interface {
	FindAll() ([]t.ExpenseOutput, error)
	FindByID(id int) (t.ExpenseToUpdate, error)
	Create(e t.ExpenseInput) (t.Expense, error)
	Update(id int, e t.ExpenseUpdate) (t.Expense, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func (s *service) FindAll() ([]t.ExpenseOutput, error) {
	return s.r.FindAll()
}

func (s *service) FindByID(id int) (t.ExpenseToUpdate, error) {
	return s.r.FindByID(id)
}

func (s *service) Create(e t.ExpenseInput) (t.Expense, error) {
	return s.r.Create(e)
}

func (s *service) Update(id int, e t.ExpenseUpdate) (t.Expense, error) {
	return s.r.Update(id, e)
}

func (s *service) Delete(id int) error {
	return s.r.Delete(id)
}

func NewService(r Repository) Service {
	return &service{r}
}
