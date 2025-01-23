package repository

import (
	"easydbTools/internal/domain/mysql/model"
)

// DatabaseRepository interface is used to define database operations
type DatabaseRepository interface {
	// GetAll get all database by dataSourceId
	GetAll(dataSourceId string) ([]model.Database, error)

	// Create a database
	Create(database model.Database) error

	// Drop a database
	Drop(database model.Database) error
}
