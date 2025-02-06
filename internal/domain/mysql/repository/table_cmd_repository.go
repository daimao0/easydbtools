package repository

import (
	"easydbTools/internal/domain/mysql/model"
)

// TableCmdRepository table command repository for cqrs
type TableCmdRepository interface {
	// CreateTable create table
	CreateTable(table *model.Table) error

	// DropTable drop table
	DropTable(table model.Table) error
}
