package controller

import (
	"easydbTools/internal/adapter/http/request"
	"easydbTools/internal/application/app"
	"easydbTools/internal/application/app/app_impl"
	"easydbTools/internal/application/cmd"
	"easydbTools/internal/application/query"
	"easydbTools/internal/common/constant"
	"easydbTools/internal/common/easytool/common/resp"
	"easydbTools/internal/common/easytool/convert"
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
	tableQuery := query.TableQuery{DataSourceId: header, DatabaseName: databaseName, TableName: tableName}
	tableDTO, err := t.tableApp.GetTable(tableQuery)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(tableDTO))
}

// CreateTable to create table
func (t *TableController) CreateTable(c *gin.Context) {
	databaseName := c.Param("databaseName")
	header := c.GetHeader(constant.XDataSourceId)
	req := &request.TableCreateRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, resp.InvalidParam(err.Error()))
		return
	}
	tableCmd := &cmd.TableCreateCmd{DataSourceId: header, DatabaseName: databaseName, Name: req.Name, Desc: req.Desc}
	tableCmd.Columns = &[]cmd.ColumnCreateCmd{}
	for _, column := range *req.Columns {
		*tableCmd.Columns = append(*tableCmd.Columns, cmd.ColumnCreateCmd{
			Name:    column.Name,
			Type:    column.Type,
			Size:    column.Size,
			Points:  column.Points,
			Default: column.Default,
			NotNull: column.NotNull,
			Comment: column.Comment,
			Pk:      column.Pk,
		})
	}
	tableCmd.Indexes = &[]cmd.IndexCmd{}
	for _, indexRequest := range *req.Indexes {
		*tableCmd.Indexes = append(*tableCmd.Indexes, cmd.IndexCmd{
			Name:       indexRequest.Name,
			Unique:     indexRequest.Unique,
			ColumnName: indexRequest.ColumnName,
			Comment:    indexRequest.Comment,
		})
	}
	err = t.tableApp.CreateTable(tableCmd)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(tableCmd))
}

// DropTable to drop table
func (t *TableController) DropTable(c *gin.Context) {
	databaseName := c.Param("databaseName")
	tableName := c.Param("tableName")
	header := c.GetHeader(constant.XDataSourceId)
	err := t.tableApp.DropTable(cmd.TableDropCmd{DataSourceId: header, DatabaseName: databaseName, Name: tableName})
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

// PageTableData to page query data
func (t *TableController) PageTableData(c *gin.Context) {
	databaseName := c.Param("databaseName")
	tableName := c.Param("tableName")
	pageNO := c.Param("pageNO")
	pageSize := c.Param("pageSize")
	dataSourceId := c.GetHeader(constant.XDataSourceId)
	dataQuery := query.TableDataQuery{TableQuery: &query.TableQuery{DataSourceId: dataSourceId, DatabaseName: databaseName, TableName: tableName}, PageNo: convert.ToInt(pageNO), PageSize: convert.ToInt(pageSize)}
	t.tableApp.PageTableData(&dataQuery)
}
