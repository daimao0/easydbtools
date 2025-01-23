package intercpter

import (
	"easydbTools/internal/common/easytool/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// DataSourceInterceptor the web request could be carried the data source id from the request header
func DataSourceInterceptor(c *gin.Context) {
	// pass the uri
	if strings.HasPrefix(c.Request.URL.Path, "/api/v1/datasource") {
		c.Next()
		return
	}
	dataSourceId := c.GetHeader("x-datasource-id")
	if dataSourceId == "" {
		c.JSON(http.StatusUnauthorized, resp.DataSourceIsNotExist)
		c.Abort()
		return
	}
	c.Next()
}
