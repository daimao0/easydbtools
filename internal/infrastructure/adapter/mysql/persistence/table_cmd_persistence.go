package persistence

import (
	"easydbTools/internal/common/enum"
	"easydbTools/internal/domain/mysql/model"
	"easydbTools/internal/domain/mysql/repository"
	"fmt"
	"strings"
)

// TableCmdPersistence is the TableCmdRepository interface implementation
type TableCmdPersistence struct {
	dataSourceRepository repository.DataSourceRepository
}

func NewTableCmdPersistence() *TableCmdPersistence {
	return &TableCmdPersistence{
		dataSourceRepository: NewDataSourceRepositoryImpl(),
	}
}

// CreateTable creates a table
func (t *TableCmdPersistence) CreateTable(table *model.Table) error {
	connect := t.dataSourceRepository.ConnectById(table.Database.DataSourceId)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s.%s (", table.Database.Name, table.Name))
	// build columns
	for _, column := range *table.Columns {
		builder.WriteString(fmt.Sprintf("%s %s", column.Name, column.Type))
		columnType, _ := enum.GetColumnType(column.Type)
		if column.Size > 0 {
			builder.WriteString(fmt.Sprintf("(%d)", column.Size))
		} else if columnType == enum.VARCHAR {
			builder.WriteString(fmt.Sprintf("(255)"))
		}
		if column.Points != "" && column.Size == 0 {
			builder.WriteString(fmt.Sprintf("(%s)", column.Points))
		}
		if column.NotNull {
			builder.WriteString(" NOT NULL")
		}
		if column.Default != "" {
			builder.WriteString(fmt.Sprintf(" DEFAULT '%s'", column.Default))
		}
		if column.Comment != "" {
			builder.WriteString(fmt.Sprintf(" COMMENT '%s'", column.Comment))
		}
		builder.WriteString(",")
	}
	// build index
	for _, index := range *table.Indexes {
		commas, err := index.ToStringColumnNames()
		if err != nil {
			return err
		}
		builder.WriteString(fmt.Sprintf("INDEX %s (%s) COMMENT '%s',", index.Name, commas, index.Comment))
	}
	// build pk
	if pk, err := table.GetPK(); err == nil {
		builder.WriteString(fmt.Sprintf("PRIMARY KEY (%s)", pk.Name))
	}
	// build table desc
	builder.WriteString(fmt.Sprintf(") COMMENT='%s';", table.Desc))
	_, err := connect.Exec(builder.String())
	return err
}

// DropTable drop table
func (t *TableCmdPersistence) DropTable(table model.Table) error {

	connect := t.dataSourceRepository.ConnectById(table.Database.DataSourceId)
	sql := fmt.Sprintf("DROP TABLE %s.%s;", table.Database.Name, table.Name)
	_, err := connect.Exec(sql)
	return err
}
