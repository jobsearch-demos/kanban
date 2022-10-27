package filter

type UserFilter struct {
	Offset   uint64   `json:"offset" lookup:"offset" operator:"eq" validate:"omitempty,gt=0" default:"0"`
	Limit    uint64   `json:"limit" lookup:"limit" operator:"eq" validate:"omitempty,gt=0" default:"10"`
	ID       int      `json:"id" lookup:"id" operator:"eq" validate:"omitempty,gt=0"`
	Username string   `json:"username" validate:"required" lookup:"username" operator:"icontains" validate:"omitempty" default:""`
	Email    string   `json:"email" lookup:"email" operator:"icontains" default:"" validate:"omitempty,email"`
	Groups   []string `json:"groups" lookup:"groups__id" operator:"in" relation:"groups" validate:"omitempty" default:""`
}
