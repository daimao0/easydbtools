package convert

import (
	"easydbTools/internal/adapter/http/request"
	"easydbTools/internal/application/cmd"
	"easydbTools/internal/application/dto"
	"easydbTools/internal/domain/mysql/model"
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
	return model.Database{DataSourceId: cmd.DataSourceId, Name: cmd.Name}
}

// DatabaseDropCmdToDatabase convert DatabaseDropCmd to Database
func DatabaseDropCmdToDatabase(cmd cmd.DatabaseDropCmd) model.Database {
	return model.Database{DataSourceId: cmd.DataSourceId, Name: cmd.Name}
}
