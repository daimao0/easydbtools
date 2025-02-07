package app_impl

import (
	"easydbTools/internal/application/cmd"
	"easydbTools/internal/application/dto"
	"easydbTools/internal/application/query"
	"easydbTools/internal/common/easytool/common/page"
	"easydbTools/internal/domain/mysql/model"
	"easydbTools/internal/domain/mysql/repository"
	"easydbTools/internal/domain/mysql/repository/params"
	"easydbTools/internal/infrastructure/adapter/mysql/persistence"
	"easydbTools/internal/infrastructure/convert"
)

type TableAppImpl struct {
	tableQueryRepository repository.TableQueryRepository
	tableCmdRepository   repository.TableCmdRepository
}

func NewTableAppImpl() *TableAppImpl {
	return &TableAppImpl{
		tableQueryRepository: persistence.NewTableQueryRepositoryImpl(),
		tableCmdRepository:   persistence.NewTableCmdPersistence(),
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
	queryParams := params.TableQueryParams{DataSourceId: tableQuery.DataSourceId, DatabaseName: tableQuery.DatabaseName, TableName: tableQuery.TableName}
	table, _ := t.tableQueryRepository.GetTable(&queryParams)
	return dto.TableDTO{
		Name:      table.Name,
		Desc:      table.Desc,
		CreateSQL: table.CreateSQL,
		Database:  &dto.DatabaseDTO{Name: table.Database.Name},
		Columns:   convert.ColumnsToColumnDTOs(table.Columns),
		Indexes:   convert.IndexesToIndexDTOs(table.Indexes),
	}, nil
}

// CreateTable creates table
func (t *TableAppImpl) CreateTable(tableCreateCmd *cmd.TableCreateCmd) error {
	columns := make([]model.Column, 0)
	for _, createCmd := range *tableCreateCmd.Columns {
		columns = append(columns, model.Column{
			Name:    createCmd.Name,
			Type:    createCmd.Type,
			Size:    createCmd.Size,
			Points:  createCmd.Points,
			Default: createCmd.Default,
			NotNull: createCmd.NotNull,
			Comment: createCmd.Comment,
			Pk:      createCmd.Pk,
		})
	}
	database := model.Database{Name: tableCreateCmd.DatabaseName, DataSourceId: tableCreateCmd.DataSourceId}

	table := &model.Table{Name: tableCreateCmd.Name, Desc: tableCreateCmd.Desc, Database: &database, Columns: &columns}
	table.Indexes = &[]model.Index{}
	for _, index := range *tableCreateCmd.Indexes {
		indexColumns := make([]model.Column, 0)
		for _, col := range index.ColumnName {
			name, err := table.GetColumnByName(col)
			if err != nil {
				return err
			}
			indexColumns = append(indexColumns, *name)
		}
		*table.Indexes = append(*table.Indexes, model.Index{
			Name:    index.Name,
			Unique:  index.Unique,
			Columns: &indexColumns,
			Comment: index.Comment,
		})
	}
	err := t.tableCmdRepository.CreateTable(table)
	return err
}

// DropTable drops a table
func (t *TableAppImpl) DropTable(dropCmd cmd.TableDropCmd) error {
	err := t.tableCmdRepository.DropTable(model.Table{Name: dropCmd.Name, Database: &model.Database{Name: dropCmd.DatabaseName, DataSourceId: dropCmd.DataSourceId}})
	return err
}

// PageTableData returns a list of table data
func (t *TableAppImpl) PageTableData(tableQuery *query.TableDataQuery) *page.Page[[]map[string]interface{}] {
	q := tableQuery.TableQuery
	pageParams := params.TablePageParams{TableQueryParams: &params.TableQueryParams{DataSourceId: q.DataSourceId, DatabaseName: q.DatabaseName, TableName: q.TableName}, PageNo: tableQuery.PageNo, PageSize: tableQuery.PageSize}
	return t.tableQueryRepository.PageTableData(&pageParams)
}
