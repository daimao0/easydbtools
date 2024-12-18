package convert

import (
	"easydbTools/internal/adapter/http/request"
	"easydbTools/internal/application/cmd"
	"easydbTools/internal/application/dto"
	"easydbTools/internal/domain/database/model"
)

// DatabaseCreateRequestToDatabaseCreateCmd convert DatabaseCreateRequest to DatabaseCreateCmd
func DatabaseCreateRequestToDatabaseCreateCmd(request request.DatabaseCreateRequest) cmd.DatabaseCreateCmd {
	return cmd.DatabaseCreateCmd{Name: request.Name}
}

// DatabaseToDatabaseDTO convert Database to DatabaseDTO
func DatabaseToDatabaseDTO(database model.Database) dto.DatabaseDTO {
	return dto.DatabaseDTO{Name: database.Name}
}

// DatabaseCreateCmdToDatabase convert DatabaseCreateCmd to Database
func DatabaseCreateCmdToDatabase(cmd cmd.DatabaseCreateCmd) model.Database {
	return model.Database{Name: cmd.Name}
}

// DatabaseDropCmdToDatabase convert DatabaseDropCmd to Database
func DatabaseDropCmdToDatabase(cmd cmd.DatabaseDropCmd) model.Database {
	return model.Database{Name: cmd.Name}
}

// DatabaseDropRequestToDatabaseDropCmd convert DatabaseDropRequest to DatabaseDropCmd
func DatabaseDropRequestToDatabaseDropCmd(dropRequest request.DatabaseDropRequest) cmd.DatabaseDropCmd {
	return cmd.DatabaseDropCmd{
		Name: dropRequest.Name,
	}
}
