package persistence

import (
	"easydbTools/internal/common/error_code"
	"easydbTools/internal/domain/mysql/model"
	"easydbTools/internal/domain/mysql/repository"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// DatabaseRepositoryImpl implements the Repository interface
type DatabaseRepositoryImpl struct {
	datasourceRepository repository.DataSourceRepository
}

// NewDatabaseRepositoryImpl returns a new instance of RepositoryImpl
func NewDatabaseRepositoryImpl() *DatabaseRepositoryImpl {
	return &DatabaseRepositoryImpl{
		datasourceRepository: NewDataSourceRepositoryImpl(),
	}
}

// GetAll gets all records from the database
func (r *DatabaseRepositoryImpl) GetAll(dataSourceId string) ([]model.Database, error) {
	connect := r.datasourceRepository.ConnectById(dataSourceId)
	if connect == nil {
		return nil, error_code.DataSourceConnectNotExists.Err
	}
	query, err := connect.Query("SHOW DATABASES")
	defer query.Close()
	// defer the close till after the main query completes
	if err != nil {
		return nil, err
	}
	var databases []model.Database
	for query.Next() {
		var databaseName string
		_ = query.Scan(&databaseName)
		databases = append(databases, model.Database{Name: databaseName})
	}
	return databases, nil
}

// Create a database
func (r *DatabaseRepositoryImpl) Create(database model.Database) error {
	connect := r.datasourceRepository.ConnectById(database.DataSourceId)
	_, err := connect.Exec(fmt.Sprintf("CREATE DATABASE %s DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;", database.Name))
	return err
}

// Drop a database
func (r *DatabaseRepositoryImpl) Drop(database model.Database) error {
	connect := r.datasourceRepository.ConnectById(database.DataSourceId)
	_, err := connect.Exec(fmt.Sprintf("DROP DATABASE %s;", database.Name))
	return err
}
