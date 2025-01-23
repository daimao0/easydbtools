package query

// TableNamesAppQuery represents a query for table app layer service
type TableNamesAppQuery struct {
	//DataSourceId is the id of the data source
	DataSourceId string

	//DatabaseName is the name of the database
	DatabaseName string
}

// TableQuery represents a query for table service
type TableQuery struct {
	//DataSourceId is the id of the data source
	DataSourceId string

	//DatabaseName is the name of the database
	DatabaseName string

	//TableName is the name of the table
	TableName string
}
