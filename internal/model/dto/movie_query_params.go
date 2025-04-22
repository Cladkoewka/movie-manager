package dto

type MovieQueryParams struct {
	Search   string   `json:"search,omitempty"` // Search term for movie title
	Genre    string   `json:"genre,omitempty"`
	Language string   `json:"language,omitempty"`
	Rating   *float64 `json:"rating,omitempty"`
	SortBy   string   `json:"sortBy,omitempty"`
	OrderBy  string   `json:"orderBy,omitempty"`
	Page     int      `json:"page,omitempty"`
	PageSize int      `json:"pageSize,omitempty"`
}
