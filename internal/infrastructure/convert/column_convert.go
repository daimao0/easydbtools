package convert

import (
	"easydbTools/internal/application/dto"
	"easydbTools/internal/domain/mysql/model"
)

func ColumnToColumnDTO(column *model.Column) *dto.ColumnDTO {
	return &dto.ColumnDTO{
		Name:    column.Name,
		Type:    column.Type,
		Size:    column.Size,
		Points:  column.Points,
		Default: column.Default,
		NotNull: column.NotNull,
		Comment: column.Comment,
		Pk:      column.Pk,
	}
}

func ColumnsToColumnDTOs(columns *[]model.Column) *[]dto.ColumnDTO {
	columnDTOs := make([]dto.ColumnDTO, 0)
	for _, column := range *columns {
		columnDTOs = append(columnDTOs, *ColumnToColumnDTO(&column))
	}
	return &columnDTOs
}
