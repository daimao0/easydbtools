package repository

import (
	"easydbTools/internal/domain/mysql/model"
	"easydbTools/internal/domain/mysql/repository/params"
)

// TableQueryRepository is used to get table information from the database
type TableQueryRepository interface {
	// ListTableNames is used to get all table names from the database
	ListTableNames(database model.Database) ([]string, error)

	// GetTable is used to get table information from the database
	GetTable(params *params.TableQueryParams) (*model.Table, error)
}
