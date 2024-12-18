package controller

import (
	"easydbTools/internal/adapter/http/request"
	"easydbTools/internal/application/app"
	"easydbTools/internal/application/app/app_impl"
	"easydbTools/internal/common/easytool/resp"
	"easydbTools/internal/common/error_code"
	"easydbTools/internal/infrastructure/convert"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DataSourceController struct {
	dataSourceApp app.DataSourceApp
}

func NewDatasourceController() *DataSourceController {
	return &DataSourceController{
		dataSourceApp: app_impl.NewDatasourceAppImpl(),
	}
}

// TestConnect datasource
func (dataSourceController *DataSourceController) TestConnect(c *gin.Context) {
	dataSourceConnectRequest := request.DataSourceConnectRequest{}
	_ = c.BindJSON(&dataSourceConnectRequest)
	err := dataSourceController.dataSourceApp.TestConnect(convert.DataSourceConnectRequestToDatasourceConnectCmd(dataSourceConnectRequest))
	if err != nil {
		c.JSON(http.StatusOK, resp.SystemError(error_code.DATA_SOURCE_CONNECT_ERROR))
		return
	}
	c.JSON(http.StatusOK, resp.SUCCESS)
}

// Connect datasource and register to memory
func (dataSourceController *DataSourceController) Connect(c *gin.Context) {
	dataSourceConnectRequest := request.DataSourceConnectRequest{}
	_ = c.BindJSON(&dataSourceConnectRequest)
	err := dataSourceController.dataSourceApp.Connect(convert.DataSourceConnectRequestToDatasourceConnectCmd(dataSourceConnectRequest))
	if err != nil {
		c.JSON(http.StatusOK, resp.SystemError(error_code.DATA_SOURCE_CONNECT_ERROR))
		return
	}
	c.JSON(http.StatusOK, resp.SUCCESS)
}
