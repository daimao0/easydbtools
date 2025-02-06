package model

import "easydbTools/internal/common/error_code"

// Table is the table model
type Table struct {

	// Name is the table name
	Name string

	// Desc is the table desc
	Desc string

	// CreateSQL is the create table sql
	CreateSQL string

	// Database is the database model
	Database *Database

	// Columns is the table columns
	Columns *[]Column

	// Indexes is the table indexes
	Indexes *[]Index
}

func (t *Table) GetPK() (*Column, error) {
	for _, column := range *t.Columns {
		if column.Pk {
			return &column, nil
		}
	}
	return nil, error_code.TablePkNotExists
}

func (t *Table) GetColumnByName(name string) (*Column, error) {
	for _, column := range *t.Columns {
		if column.Name == name {
			return &column, nil
		}
	}
	return nil, error_code.TablePkNotExists
}
