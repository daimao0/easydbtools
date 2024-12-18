package app

import (
	"easydbTools/internal/application/cmd"
)

// DataSourceApp is the interface for app layer
type DataSourceApp interface {

	// TestConnect the data source
	TestConnect(cmd cmd.DataSourceConnectCmd) error

	// Connect the data source and cache in memory
	Connect(cmd cmd.DataSourceConnectCmd) error
}
