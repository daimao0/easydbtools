package controller

import (
	"easydbTools/internal/application/app"
	"easydbTools/internal/application/app/app_impl"
	"easydbTools/internal/application/query"
	"easydbTools/internal/common/constant"
	"easydbTools/internal/common/easytool/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TableController struct {
	tableApp app.TableApp
}

func NewTableController() *TableController {
	return &TableController{
		tableApp: app_impl.NewTableAppImpl(),
	}
}

// ListTables represents get all tables name from the database
// get /api/v1/databases/:databaseName/tables
func (t *TableController) ListTables(c *gin.Context) {
	databaseName := c.Param("databaseName")
	dataSourceId := c.GetHeader(constant.XDataSourceId)
	// construct the query
	appQuery := query.TableNamesAppQuery{DataSourceId: dataSourceId, DatabaseName: databaseName}
	names, err := t.tableApp.ListTableNames(appQuery)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(names))
}

// GetTable represents get table detail from the database by table name
func (t *TableController) GetTable(c *gin.Context) {
	databaseName := c.Param("databaseName")
	tableName := c.Param("tableName")
	header := c.GetHeader(constant.XDataSourceId)
	tableQuery := query.TableQuery{DataSourceId: databaseName, DatabaseName: tableName, TableName: header}
	column, err := t.tableApp.GetTable(tableQuery)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(column))
}
