package types


type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}