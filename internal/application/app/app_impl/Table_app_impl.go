package app_impl

import (
	"easydbTools/internal/application/dto"
	"easydbTools/internal/application/query"
	"easydbTools/internal/domain/mysql/model"
	"easydbTools/internal/domain/mysql/repository"
	"easydbTools/internal/infrastructure/adapter/mysql/persistence"
)

type TableAppImpl struct {
	tableQueryRepository repository.TableQueryRepository
}

func NewTableAppImpl() *TableAppImpl {
	return &TableAppImpl{
		tableQueryRepository: persistence.NewTableQueryRepositoryImpl(),
	}
}

// ListTableNames is show tables names from the database
func (t *TableAppImpl) ListTableNames(query query.TableNamesAppQuery) ([]string, error) {
	// convert app to query
	database := model.Database{DataSourceId: query.DataSourceId, Name: query.DatabaseName}
	//  query table names from table repository
	names, err := t.tableQueryRepository.ListTableNames(database)
	if err != nil {
		return nil, err
	}
	return names, nil
}

// GetTable is show table  from the database
func (t *TableAppImpl) GetTable(tableQuery query.TableQuery) (dto.TableDTO, error) {
	//t.tableQueryRepository.GetTable(tableQuery)
	return dto.TableDTO{}, nil
}
