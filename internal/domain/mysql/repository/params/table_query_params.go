package params

// TableQueryParams is the parameters for querying a table
type TableQueryParams struct {
	//DataSourceId is the id of the data source
	DataSourceId string

	//DatabaseName is the name of the database
	DatabaseName string

	//TableName is the name of the table
	TableName string
}
