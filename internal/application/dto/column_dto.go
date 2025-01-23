package dto

// ColumnDTO is the data transfer object for a column
type ColumnDTO struct {
	// Name is the name of the column
	Name string
	// Type is the type of the column
	Type string
	// IsPrimary is true if the column is a primary key
	IsPrimary bool
	// IsNullable is true if the column can be null
	IsNullable bool
	// IsUnique is true if the column is unique
	IsUnique bool
	// IsAuto is true if the column is auto-incrementing
	IsAuto bool
	// Default is the default value of the column
	Default string
	// Comment is the comment of the column
	Comment string
}
