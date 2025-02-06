package cmd

// IndexCmd is a common struct for index command
type IndexCmd struct {

	// Name  the name of the index
	Name string

	// ColumnName lists the columns that will be used to create the index. These columns' data will be utilized to speed up queries.
	ColumnName []string

	// Unique indicates whether the index should enforce uniqueness. If set to true, the index ensures that all index key values are unique.
	Unique bool

	// Comment allows adding a descriptive note to the index. This can help explain the purpose of the index or provide usage considerations.
	Comment string
}
