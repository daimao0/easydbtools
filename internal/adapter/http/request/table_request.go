package request

// TableCreateRequest create table request
type TableCreateRequest struct {

	// Name table name
	Name string `json:"name"`

	// Desc table desc
	Desc string `json:"desc"`

	// Columns create table columns
	Columns *[]ColumnCreateRequest `json:"columns"`

	// Indexes create table indexes
	Indexes *[]IndexRequest `json:"indexes"`
}
