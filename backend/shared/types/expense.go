package types

type Expense struct {
	ID          int     `json:"id" db:"id"`
	Description string  `json:"description" db:"description"`
	Amount      float64 `json:"amount" db:"amount"`
	ProfileID   int     `json:"profile_id" db:"profile_id"`
}

type ExpenseInput struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	CategoryID  int     `json:"category_id"`
	ProfileID   int     `json:"profile_id" db:"profile_id"`
}

type ExpenseUpdate struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	CategoryID  int     `json:"category_id"`
}

type ExpenseToUpdate struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category" db:"category"`
}

type ExpenseOutput struct {
	ID          int     `json:"id" db:"id"`
	Description string  `json:"description" db:"description"`
	Amount      float64 `json:"amount" db:"amount"`
	Category    string  `json:"category" db:"category"`
	SubTotal    float64 `json:"sub_total" db:"sub_total"`
}
