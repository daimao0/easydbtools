package app_impl

import (
	"easydbTools/internal/application/cmd"
	"easydbTools/internal/application/dto"
	"easydbTools/internal/domain/database/repository"
	"easydbTools/internal/infrastructure/adapter/persistence/database/repository_impl"
	"easydbTools/internal/infrastructure/convert"
)

// DatabaseAppImpl layer
type DatabaseAppImpl struct {
	databaseRepository repository.DatabaseRepository
}

// NewDatabaseAppImpl new DatabaseApp
func NewDatabaseAppImpl() *DatabaseAppImpl {
	return &DatabaseAppImpl{
		databaseRepository: repository_impl.NewDatabaseRepositoryImpl(),
	}
}

// List all databases
func (app *DatabaseAppImpl) List(dataSourceId string) ([]dto.DatabaseDTO, error) {
	databases, err := app.databaseRepository.GetAll(dataSourceId)
	if err != nil {
		return nil, err
	}
	databaseDTOs := make([]dto.DatabaseDTO, len(databases))
	for i := range databases {
		databaseDTOs[i] = convert.DatabaseToDatabaseDTO(databases[i])
	}
	return databaseDTOs, nil
}

// Create a database
func (app *DatabaseAppImpl) Create(cmd cmd.DatabaseCreateCmd) error {
	database := convert.DatabaseCreateCmdToDatabase(cmd)
	err := app.databaseRepository.Create(database)
	return err
}

// Drop a database
func (app *DatabaseAppImpl) Drop(cmd cmd.DatabaseDropCmd) error {
	database := convert.DatabaseDropCmdToDatabase(cmd)
	err := app.databaseRepository.Drop(database)
	return err
}
