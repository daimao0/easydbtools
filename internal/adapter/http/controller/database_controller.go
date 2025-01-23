package controller

import (
	"easydbTools/internal/adapter/http/request"
	"easydbTools/internal/application/app"
	"easydbTools/internal/application/app/app_impl"
	"easydbTools/internal/application/cmd"
	"easydbTools/internal/common/constant"
	"easydbTools/internal/common/easytool/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DatabaseController struct {
	databaseApp app.DatabaseApp
}

func NewDatabaseController() *DatabaseController {
	return &DatabaseController{
		databaseApp: app_impl.NewDatabaseAppImpl(),
	}
}

// RegisterDataSource register data source to memory
func (databaseController *DatabaseController) RegisterDataSource(c *gin.Context) {

}

// List show database
func (databaseController *DatabaseController) List(c *gin.Context) {
	dataSourceId := c.GetHeader("x-datasource-id")
	databases, err := databaseController.databaseApp.List(dataSourceId)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(databases))
}

// Create database
func (databaseController *DatabaseController) Create(c *gin.Context) {
	createRequest := request.DatabaseCreateRequest{}
	err := c.BindJSON(&createRequest)
	if err != nil {
		c.JSON(http.StatusOK, resp.INVALID_PARAM)
		return
	}
	// convert request to cmd
	createCmd := cmd.DatabaseCreateCmd{DataSourceId: c.GetHeader(constant.XDataSourceId), Name: createRequest.Name}
	// create a database
	err = databaseController.databaseApp.Create(createCmd)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

// Drop database
func (databaseController *DatabaseController) Drop(c *gin.Context) {
	//convert request to cmd
	dropCmd := cmd.DatabaseDropCmd{DataSourceId: c.GetHeader(constant.XDataSourceId), Name: c.Param("name")}
	err := databaseController.databaseApp.Drop(dropCmd)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}
