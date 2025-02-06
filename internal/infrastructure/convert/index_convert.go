package convert

import (
	"easydbTools/internal/application/dto"
	"easydbTools/internal/domain/mysql/model"
)

func IndexToIndexDTO(index *model.Index) *dto.IndexDTO {
	return &dto.IndexDTO{
		Name:    index.Name,
		Unique:  index.Unique,
		Columns: ColumnsToColumnDTOs(index.Columns),
		Comment: index.Comment,
	}
}

func IndexesToIndexDTOs(indexes *[]model.Index) *[]dto.IndexDTO {
	indexDTOs := make([]dto.IndexDTO, 0)
	for _, index := range *indexes {
		indexDTOs = append(indexDTOs, *IndexToIndexDTO(&index))
	}
	return &indexDTOs
}
