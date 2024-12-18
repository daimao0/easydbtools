package request

type DatabaseDatasourceRequest struct {
	// data source connect request
	DataSourceConnectRequest DataSourceConnectRequest `json:"dataSourceConnectRequest"`
	// database name
	Name string `json:"name"`
	// charset
	Charset string `json:"charset"`
}

// DatabaseCreateRequest is the request body for creating a database
type DatabaseCreateRequest struct {
	// Name of the database
	Name string `json:"name"`
	// Charset of the database
	Charset string `json:"charset"`
}

// DatabaseDropRequest is the request body for dropping a database
type DatabaseDropRequest struct {
	// Name of the database
	Name string `json:"name"`
}
