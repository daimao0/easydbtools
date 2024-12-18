package convert

import (
	"easydbTools/internal/adapter/http/request"
	"easydbTools/internal/application/cmd"
	"easydbTools/internal/common/enum"
	"easydbTools/internal/domain/datasource/model"
)

// DatasourceConnectCmdToDatasource  convert datasourceConnectCmd to datasource
func DatasourceConnectCmdToDatasource(cmd cmd.DataSourceConnectCmd) model.DataSource {
	return model.DataSource{
		Id:         cmd.Id,
		DriverName: cmd.DriverName,
		Name:       cmd.Name,
		Address:    cmd.Address,
		Username:   cmd.Username,
		Password:   cmd.Password,
	}
}

// DataSourceConnectRequestToDatasourceConnectCmd convert datasourceConnectRequest to datasourceConnectCmd
func DataSourceConnectRequestToDatasourceConnectCmd(request request.DataSourceConnectRequest) cmd.DataSourceConnectCmd {
	name, err := enum.GetDriverName(request.DriverName)
	if err != nil {
		name = ""
	}
	return cmd.DataSourceConnectCmd{
		Id:         request.Id,
		DriverName: name,
		Name:       request.Name,
		Address:    request.Address,
		Username:   request.Username,
		Password:   request.Password,
	}
}
