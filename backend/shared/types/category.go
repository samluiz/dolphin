package types

type Category struct {
	ID          int    `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type CategoryInput struct {
	Description string `json:"description"`
}
