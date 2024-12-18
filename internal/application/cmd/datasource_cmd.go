package cmd

import "easydbTools/internal/common/enum"

// DataSourceConnectCmd is the datasource connect command
type DataSourceConnectCmd struct {
	Id string

	DriverName enum.DriverName

	Name string

	Address string

	Username string

	Password string
}
