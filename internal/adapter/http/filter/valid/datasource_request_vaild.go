package valid

import (
	"bytes"
	"easydbTools/internal/adapter/http/request"
	"easydbTools/internal/common/easytool/common/resp"
	"easydbTools/internal/common/easytool/util/str_util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func DataSourceConnectRequestValid(request *request.DataSourceConnectRequest) gin.HandlerFunc {
	return func(c *gin.Context) {

		data, _ := c.GetRawData()
		err := json.Unmarshal(data, request)
		if err != nil {
			c.JSON(http.StatusOK, resp.INVALID_PARAM)
			c.Abort()
			return
		}
		if str_util.IsBlank(request.Name) || str_util.IsBlank(request.Address) || str_util.IsBlank(request.Username) || str_util.IsBlank(request.Password) {
			c.JSON(http.StatusOK, resp.INVALID_PARAM)
			c.Abort()
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
		c.Next()
	}
}
