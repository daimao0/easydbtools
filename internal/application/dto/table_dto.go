package dto

// TableDTO is the app layer representation of a table
type TableDTO struct {

	// Name is the name of the table
	Name string

	// Columns is the list of columns in the table
	Columns []ColumnDTO
}

// GetPK gets the primary key of the table
// If the table has no primary key, it returns nil
func (t *TableDTO) GetPK() *ColumnDTO {
	for _, column := range t.Columns {
		if column.IsPrimary {
			return &column
		}
	}
	return nil
}
