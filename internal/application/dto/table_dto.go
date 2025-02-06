package dto

// TableDTO is the app layer representation of a table
type TableDTO struct {

	// Name is the name of the table
	Name string `json:"name"`

	// Desc is the table desc
	Desc string `json:"desc"`

	// CreateSQL is the create table sql
	CreateSQL string `json:"createSQL"`

	// Database is the database model
	Database *DatabaseDTO `json:"database"`

	// Columns is the list of columns in the table
	Columns *[]ColumnDTO `json:"columns"`

	// Indexes is the table indexes
	Indexes *[]IndexDTO `json:"indexes"`
}

// GetPK gets the primary key of the table
// If the table has no primary key, it returns nil
func (t *TableDTO) GetPK() *ColumnDTO {
	for _, column := range *t.Columns {
		if column.Pk {
			return &column
		}
	}
	return nil
}
