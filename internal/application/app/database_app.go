package app

import (
	"easydbTools/internal/application/cmd"
	"easydbTools/internal/application/dto"
)

// DatabaseApp layer
type DatabaseApp interface {

	// List all databases
	List(dataSourceId string) ([]dto.DatabaseDTO, error)

	// Create a database
	Create(cmd cmd.DatabaseCreateCmd) error

	// Drop a database
	Drop(dropCmd cmd.DatabaseDropCmd) error
}
