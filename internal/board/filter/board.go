package filter

type BoardFilter struct {
	Offset      uint64   `json:"offset,omitempty" validate:"omitempty,gt=0" default:"0"`
	Limit       uint64   `json:"limit,omitempty" validate:"omitempty,gt=0" default:"10"`
	ID          string   `json:"id,omitempty" validate:"omitempty"`
	Name        string   `json:"name,omitempty" validate:"omitempty"`
	Description string   `json:"description,omitempty" validate:"omitempty"`
	Members     []string `json:"members,omitempty" validate:"omitempty" lookup:"members__id" operator:"in" relation:"members"`
	Columns     []string `json:"columns,omitempty" validate:"omitempty" lookup:"columns__id" operator:"in" relation:"columns"`
}