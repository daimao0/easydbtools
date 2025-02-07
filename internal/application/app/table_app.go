package app

import (
	"easydbTools/internal/application/cmd"
	"easydbTools/internal/application/dto"
	"easydbTools/internal/application/query"
	"easydbTools/internal/common/easytool/common/page"
)

// TableApp is the interface for the table app layer
type TableApp interface {
	// ListTableNames returns a list of tableNames
	ListTableNames(query query.TableNamesAppQuery) ([]string, error)

	// GetTable returns a list of columns
	GetTable(tableQuery query.TableQuery) (dto.TableDTO, error)

	// CreateTable creates a table
	CreateTable(tableCreateCmd *cmd.TableCreateCmd) error

	// DropTable drops a table
	DropTable(dropCmd cmd.TableDropCmd) error

	// PageTableData returns a list of table data
	PageTableData(tableQuery *query.TableDataQuery) *page.Page[[]map[string]interface{}]
}
