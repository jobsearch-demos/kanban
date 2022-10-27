package filter

type ProfileFilter struct {
	Offset      uint64 `json:"offset,omitempty" default:"0" validate:"omitempty,gt=0"`
	Limit       uint64 `json:"limit,omitempty" default:"10" validate:"omitempty,gt=0"`
	ID          string `json:"id,omitempty" validate:"omitempty"`
	FirstName   string `json:"first_name,omitempty" validate:"omitempty"`
	LastName    string `json:"last_name,omitempty" validate:"omitempty"`
	Email       string `json:"email,omitempty" validate:"omitempty,email"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"omitempty,numeric"`
}
