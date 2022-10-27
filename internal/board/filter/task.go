package filter

type TaskFilter struct {
	Offset       uint64 `json:"offset" validate:"omitempty,gt=0" default:"0"`
	Limit        uint64 `json:"limit" validate:"omitempty,gt=0" default:"10"`
	ID           string `json:"id" validate:"omitempty"`
	BoardID      string `json:"board_id" validate:"omitempty"`
	ColumnID     string `json:"column_id" validate:"omitempty"`
	Name         string `json:"name" validate:"omitempty"`
	Description  string `json:"description" validate:"omitempty"`
	Order        uint64 `json:"order" validate:"omitempty,number"`
	CreatedAtGTE string `json:"created_at__gte" validate:"omitempty,datetime"`
	CreatedAtLTE string `json:"created_at__lte" validate:"omitempty,datetime"`
	UpdatedAtGTE string `json:"updated_at__gte" validate:"omitempty,datetime"`
	UpdatedAtLTE string `json:"updated_at__lte" validate:"omitempty,datetime"`
}
