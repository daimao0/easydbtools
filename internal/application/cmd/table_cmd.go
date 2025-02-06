package cmd

// TableCreateCmd struct
type TableCreateCmd struct {

	// DataSourceId used to connect to database
	DataSourceId string

	// DatabaseName which table belongs the database
	DatabaseName string

	// Name of table
	Name string

	// Desc of table
	Desc string

	// Columns of table
	Columns *[]ColumnCreateCmd

	// Indexes of table
	Indexes *[]IndexCmd
}

type TableDropCmd struct {

	// DataSourceId used to connect to database
	DataSourceId string

	// DatabaseName which table belongs the database
	DatabaseName string

	// Name of table
	Name string
}
