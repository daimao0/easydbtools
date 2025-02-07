package page

// Page struct
type Page[T any] struct {
	// PageNo is the page number
	PageNo int `json:"pageNo"`
	// PageSize is the page size
	PageSize int `json:"pageSize"`
	// Total is the total number of items
	Total int `json:"total"`
	// Data is the page data
	Data []T `json:"data"`
}
