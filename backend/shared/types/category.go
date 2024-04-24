package types

type Category struct {
	ID          int    `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
}

type CategoryInput struct {
	Description string `json:"description"`
}
