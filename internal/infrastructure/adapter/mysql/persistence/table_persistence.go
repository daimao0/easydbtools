package persistence

import (
	"easydbTools/internal/domain/mysql/model"
	"easydbTools/internal/domain/mysql/repository"
	"easydbTools/internal/domain/mysql/repository/params"
	"fmt"
)

type TableQueryRepositoryImpl struct {
	datasourceRepository repository.DataSourceRepository
}

func NewTableQueryRepositoryImpl() *TableQueryRepositoryImpl {
	return &TableQueryRepositoryImpl{
		datasourceRepository: NewDataSourceRepositoryImpl(),
	}
}

// ListTableNames is used to get all table names from the database
func (t *TableQueryRepositoryImpl) ListTableNames(database model.Database) ([]string, error) {
	// get the connection
	connect := t.datasourceRepository.ConnectById(database.DataSourceId)
	// build the sql
	executeSql := fmt.Sprintf("show tables from %s;", database.Name)
	rows, err := connect.Query(executeSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// scan the result
	var tableNames []string
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tableNames = append(tableNames, tableName)
	}
	return tableNames, nil
}

// GetTable is used to get the table information
func (t *TableQueryRepositoryImpl) GetTable(params params.TableQueryParams) (model.Table, error) {
	//connect := t.datasourceRepository.ConnectById(params.DataSourceId)
	//connect.Exec()
	//connect.Query()
	return model.Table{}, nil
}
