package dto

// IndexDTO the table index
type IndexDTO struct {

	// Name the index name
	Name string `json:"name"`

	// unique the index is unique
	Unique bool `json:"unique"`

	// Columns the index columns
	Columns *[]ColumnDTO `json:"columns"`

	//Comment the index comment
	Comment string `json:"comment"`
}
