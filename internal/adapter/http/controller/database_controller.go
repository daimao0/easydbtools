package controller

import (
	"easydbTools/internal/adapter/http/request"
	"easydbTools/internal/application/app"
	"easydbTools/internal/application/app/app_impl"
	"easydbTools/internal/common/easytool/resp"
	"easydbTools/internal/infrastructure/convert"
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
	cmd := convert.DatabaseCreateRequestToDatabaseCreateCmd(createRequest)
	// app create a database
	err = databaseController.databaseApp.Create(cmd)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

// Drop database
func (databaseController *DatabaseController) Drop(c *gin.Context) {
	dropRequest := request.DatabaseDropRequest{}
	_ = c.BindJSON(&dropRequest)
	cmd := convert.DatabaseDropRequestToDatabaseDropCmd(dropRequest)
	err := databaseController.databaseApp.Drop(cmd)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}
