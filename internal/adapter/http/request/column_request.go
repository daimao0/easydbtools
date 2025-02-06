package request

// ColumnCreateRequest create column request for create table
type ColumnCreateRequest struct {
	// Name column name
	Name string `json:"name"`
	// Type column type
	Type string `json:"type"`
	// Size column size
	Size int `json:"size"`
	// Points column points for decimal type
	Points string `json:"points"`
	// Default column default value
	Default string `json:"default"`
	// NotNull column not null
	NotNull bool `json:"notNull"`
	// Comment column comment
	Comment string `json:"comment"`
	// Pk column pk
	Pk bool `json:"pk"`
}
