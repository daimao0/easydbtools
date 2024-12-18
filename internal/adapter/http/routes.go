package http

import (
	"easydbTools/internal/adapter/http/controller"
	"easydbTools/internal/adapter/http/filter/valid"
	"easydbTools/internal/adapter/http/request"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// RegisterRoutes register all routes
func RegisterRoutes(engine *gin.Engine) {
	// new controller
	databaseController := controller.NewDatabaseController()
	datasourceController := controller.NewDatasourceController()

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
	{
		databaseGroup.POST("register", valid.DataSourceRegisterRequestValid(&request.DatabaseDatasourceRequest{}), databaseController.RegisterDataSource)
		databaseGroup.GET("/list", databaseController.List)
		databaseGroup.POST("/create", valid.DatabaseCreateRequestValid(&request.DatabaseCreateRequest{}), databaseController.Create)
		databaseGroup.DELETE("/drop", valid.DatabaseDropRequestValid(&request.DatabaseDropRequest{}), databaseController.Drop)
	}
}
