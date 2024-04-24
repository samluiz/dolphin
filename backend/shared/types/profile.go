package types

type Profile struct {
	ID       int    `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
}

type ProfileInput struct {
	Description string `json:"description"`
}