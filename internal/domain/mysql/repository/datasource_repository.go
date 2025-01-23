package repository

import (
	"database/sql"
	"easydbTools/internal/domain/mysql/model"
)

// DataSourceRepository interface
type DataSourceRepository interface {
	// TestConnect to the data source
	TestConnect(datasource model.DataSource) error

	// Connect to the data source and cache in memory
	Connect(datasource model.DataSource) (*sql.DB, error)

	// ConnectById connect to the data source by dataSourceId
	ConnectById(id string) *sql.DB
}
