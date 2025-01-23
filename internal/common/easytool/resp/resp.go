package resp

import "easydbTools/internal/common/error_code"

// Resp struct
type Resp struct {
	// Code api status
	Code int `json:"code"`
	// Message api message
	Message string `json:"message"`
	// Data api data
	Data interface{} `json:"data"`
}

var (
	// SUCCESS api status
	SUCCESS = &Resp{200, "", nil}

	// FAIL api status
	FAIL = &Resp{500, "操作失败", nil}

	// INVALID_PARAM api status
	INVALID_PARAM = &Resp{400, "参数错误", nil}

	// UNAUTHORIZED api status
	UNAUTHORIZED = &Resp{401, "未授权", nil}

	// DataSourceIsNotExist api status
	DataSourceIsNotExist = &Resp{401, "数据源不存在", nil}
)

// Fail return fail resp
func Fail(message string) Resp {
	return Resp{
		Code:    FAIL.Code,
		Message: message,
	}
}

// Success return success resp
func Success(data interface{}) Resp {
	return Resp{
		Code: SUCCESS.Code,
		Data: data,
	}
}

// Unauthorized return unauthorized resp
func Unauthorized() Resp {
	return Resp{
		Code:    UNAUTHORIZED.Code,
		Message: UNAUTHORIZED.Message,
	}
}

// InvalidParam return invalidParam resp
func InvalidParam(msg string) Resp {
	return Resp{
		Code:    INVALID_PARAM.Code,
		Message: msg,
	}
}

// SystemError return system error resp
func SystemError(errorCode *error_code.ErrorCode) Resp {
	return Resp{
		Code:    errorCode.Code,
		Message: errorCode.Msg + "\n错误信息：",
		Data:    nil,
	}
}
