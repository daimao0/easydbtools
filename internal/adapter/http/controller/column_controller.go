package controller

import (
	"easydbTools/internal/application/query"
	"easydbTools/internal/common/constant"
	"easydbTools/internal/common/easytool/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ColumnController struct {
}

func NewColumnController() *ColumnController {
	return &ColumnController{}
}

// ListColumnsByTable get the table columns
func (cc *ColumnController) ListColumnsByTable(c *gin.Context) {

	databaseName := c.Param("databaseName")
	tableName := c.Param("tableName")
	header := c.GetHeader(constant.XDataSourceId)
	tableQuery := query.TableQuery{DataSourceId: databaseName, DatabaseName: tableName, TableName: header}

	c.JSON(http.StatusOK, resp.Success(tableQuery))
}
