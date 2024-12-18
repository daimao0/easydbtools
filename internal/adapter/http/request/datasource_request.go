package request

// DataSourceConnectRequest is the request body for testing a data source connection
type DataSourceConnectRequest struct {

	// Id of the data source by web
	Id string `json:"id"`

	// DriverName of the data source
	DriverName string `json:"driverName"`

	// Name of the data source by user custom
	Name string `json:"name"`

	// Address of the data source; Example: 127.0.0.1:3306
	Address string `json:"address"`

	// Username of the data source
	Username string `json:"username"`

	// Password of the data source
	Password string `json:"password"`

	// Database of the data source
	Description string `json:"description"`
}
