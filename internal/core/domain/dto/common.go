package dto

type List[T interface{}] struct {
	From   *int      `json:"from,omitempty"`
	Limit  *int      `json:"limit,omitempty"`
	Filter *T        `json:"filter,omitempty"`
	Search *string   `json:"search,omitempty"`
	Sort   *[]string `json:"sort,omitempty"`
}
