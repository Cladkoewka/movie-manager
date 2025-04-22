package constants

var AllowedSortFields = map[string]bool{
	"title":        true,
	"rating":       true,
	"release_date": true,
}

const (
	DefaultSortBy  = "title"
	DefaultOrderBy = "asc"
	DefaultPage    = 1
	DefaultPageSize = 10
)
