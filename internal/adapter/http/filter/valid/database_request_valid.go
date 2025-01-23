package valid

import (
	"bytes"
	"easydbTools/internal/adapter/http/request"
	"easydbTools/internal/common/easytool/resp"
	"easydbTools/internal/common/easytool/util/str_util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func DataSourceRegisterRequestValid(request *request.DatabaseDatasourceRequest) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.GetRawData()
		err := json.Unmarshal(data, request)
		if err != nil {
			c.JSON(http.StatusOK, resp.INVALID_PARAM)
			c.Abort()
			return
		}
		if str_util.IsBlank(request.Name) || str_util.IsBlank(request.Charset) || str_util.IsBlank(request.DataSourceConnectRequest.DriverName) ||
			str_util.IsBlank(request.DataSourceConnectRequest.Name) || str_util.IsBlank(request.DataSourceConnectRequest.Address) ||
			str_util.IsBlank(request.DataSourceConnectRequest.Username) || str_util.IsBlank(request.DataSourceConnectRequest.Password) {
			c.JSON(http.StatusOK, resp.INVALID_PARAM)
			c.Abort()
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
		c.Next()
	}
}

// DatabaseCreateRequestValid middleware for database create request valid
func DatabaseCreateRequestValid(request *request.DatabaseCreateRequest) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.GetRawData()
		err := json.Unmarshal(data, request)
		if err != nil {
			c.JSON(http.StatusOK, resp.INVALID_PARAM)
			c.Abort()
			return
		}
		if str_util.IsBlank(request.Name) {
			c.JSON(http.StatusOK, resp.INVALID_PARAM)
			c.Abort()
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
		c.Next()
	}

}

// DatabaseDropRequestValid middleware for database drop request valid
func DatabaseDropRequestValid() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if str_util.IsBlank(name) {
			c.JSON(http.StatusOK, resp.InvalidParam("database name cannot be blank "))
			c.Abort()
			return
		}
		c.Next()
	}
}
