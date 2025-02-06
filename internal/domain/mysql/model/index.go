package model

import (
	"easydbTools/internal/common/easytool/util/slice_utils"
	"easydbTools/internal/common/error_code"
	"strings"
)

// Index the table index
type Index struct {

	// Name the index name
	Name string

	// unique the index is unique
	Unique bool

	// Columns the index columns
	Columns *[]Column

	//Comment the index comment
	Comment string
}

// ToStringColumnNames converts the index used column names
func (index *Index) ToStringColumnNames() (string, error) {
	if slice_utils.IsEmpty(index.Columns) {
		return "", error_code.IndexNotContainsColumn
	}
	builder := strings.Builder{}
	for i, column := range *index.Columns {
		builder.WriteString(column.Name)
		if i < len(*index.Columns)-1 {
			builder.WriteString(",")
		}
	}
	return builder.String(), nil
}
