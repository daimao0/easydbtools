package cmd

// ColumnCreateCmd create column cmd for create table
type ColumnCreateCmd struct {
	// Name column name
	Name string
	// Type column type
	Type string
	// Size column size
	Size int
	// Points column points for decimal type
	Points string
	// Default column default value
	Default string
	// NotNull column not null
	NotNull bool
	// Comment column comment
	Comment string
	// Pk column pk
	Pk bool
}
