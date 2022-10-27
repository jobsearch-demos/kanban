package filter

type ColumnFilter struct {
	Offset      uint64 `json:"offset,omitempty" validate:"omitempty,gt=0" default:"0"`
	Limit       uint64 `json:"limit,omitempty" validate:"omitempty,gt=0" default:"10"`
	ID          string `json:"id,omitempty" validate:"omitempty"`
	Name        string `json:"name,omitempty" validate:"omitempty"`
	Description string `json:"description,omitempty" validate:"omitempty"`
	BoardID     string `json:"board_id,omitempty" validate:"omitempty"`
}
