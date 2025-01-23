package model

// Table is the table model
type Table struct {

	// Name is the table name
	Name string

	// CreateSQL is the create table sql
	CreateSQL string

	// Database is the database model
	Database Database

	// Columns is the table columns
	Columns []Column
}
