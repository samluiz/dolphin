package types

type Earning struct {
	ID          int     `json:"id" db:"id"`
	Description string  `json:"description" db:"description"`
	Amount      float64 `json:"amount" db:"amount"`
	ProfileID   int     `json:"profile_id" db:"profile_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type EarningInput struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	ProfileID   int     `json:"profile_id" db:"profile_id"`
}

type EarningUpdate struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

type EarningToUpdate struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

type EarningOutput struct {
	ID          int     `json:"id" db:"id"`
	Description string  `json:"description" db:"description"`
	Amount      float64 `json:"amount" db:"amount"`
	SubTotal    float64 `json:"sub_total" db:"sub_total"`
}
