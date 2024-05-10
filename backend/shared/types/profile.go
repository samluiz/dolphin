package types

type Profile struct {
	ID       int    `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type ProfileInput struct {
	Description string `json:"description"`
}