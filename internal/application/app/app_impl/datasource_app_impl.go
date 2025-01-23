package app_impl

import (
	"easydbTools/internal/application/app"
	"easydbTools/internal/application/cmd"
	"easydbTools/internal/domain/mysql/repository"
	"easydbTools/internal/infrastructure/adapter/mysql/persistence"
	"easydbTools/internal/infrastructure/convert"
)

// DatasourceAppImpl is the implementation of DatasourceApp
type DatasourceAppImpl struct {
	repository.DataSourceRepository
}

func NewDatasourceAppImpl() app.DataSourceApp {
	return &DatasourceAppImpl{
		DataSourceRepository: persistence.NewDataSourceRepositoryImpl(),
	}
}

// TestConnect test connect
func (d *DatasourceAppImpl) TestConnect(cmd cmd.DataSourceConnectCmd) error {
	datasource := convert.DatasourceConnectCmdToDatasource(cmd)
	err := d.DataSourceRepository.TestConnect(datasource)
	return err
}

// Connect datasource and cache in memory
func (d *DatasourceAppImpl) Connect(cmd cmd.DataSourceConnectCmd) error {
	datasource := convert.DatasourceConnectCmdToDatasource(cmd)
	_, err := d.DataSourceRepository.Connect(datasource)
	return err
}
