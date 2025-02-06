package http

import (
	"easydbTools/internal/adapter/http/controller"
	intercpter "easydbTools/internal/adapter/http/filter/interceptor"
	"easydbTools/internal/adapter/http/filter/valid"
	"easydbTools/internal/adapter/http/request"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// RegisterRoutes register all routes
func RegisterRoutes(engine *gin.Engine) {
	// new controller
	datasourceController := controller.NewDatasourceController()
	databaseController := controller.NewDatabaseController()
	tableController := controller.NewTableController()
	columnController := controller.NewColumnController()
	// handle cors config
	config := cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有来源
		AllowMethods:     []string{"*"}, // 允许的 HTTP 方法
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Datasource-Id"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // 是否允许发送凭据
		MaxAge:           12 * time.Hour, // 预检请求的有效期
	}
	engine.Use(cors.New(config))
	engine.Use(intercpter.DataSourceInterceptor)
	// GET /api/v1/database/list
	group := engine.Group("/api/")
	v1 := group.Group("/v1")
	//datasource route group
	datasourceGroup := v1.Group("/datasource")
	{
		datasourceGroup.POST("/test-connect", valid.DataSourceConnectRequestValid(&request.DataSourceConnectRequest{}), datasourceController.TestConnect)
		datasourceGroup.POST("/connect", valid.DataSourceConnectRequestValid(&request.DataSourceConnectRequest{}), datasourceController.Connect)
	}
	// database route group
	databaseGroup := v1.Group("/database")
	// database
	{
		databaseGroup.POST("register", valid.DataSourceRegisterRequestValid(&request.DatabaseDatasourceRequest{}), databaseController.RegisterDataSource)
		databaseGroup.GET("/list", databaseController.List)
		databaseGroup.POST("/create", valid.DatabaseCreateRequestValid(&request.DatabaseCreateRequest{}), databaseController.Create)
		databaseGroup.DELETE("/drop/:name", valid.DatabaseDropRequestValid(), databaseController.Drop)
	}
	// tables
	{
		databaseGroup.GET("/:databaseName/tables", tableController.ListTables)
		databaseGroup.GET("/:databaseName/table/:tableName", tableController.GetTable)
		databaseGroup.POST("/:databaseName/table", tableController.CreateTable)
		databaseGroup.DELETE("/:databaseName/table/:tableName", tableController.DropTable)
	}
	// columns
	{
		databaseGroup.GET("/:databaseName/table/:tableName/columns", columnController.ListColumnsByTable)
	}
}
