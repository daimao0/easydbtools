package error_code

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ErrorCode system error code
type ErrorCode struct {
	Code int
	Msg  string
	Err  error
}

var (
	DataSourceConnectError     = &ErrorCode{10001, "数据源连接异常", errors.New("数据源连接异常")}
	DataSourceConnectNotExists = &ErrorCode{10002, "数据源不存在", errors.New("数据源不存在")}
	TablePkNotExists           = &ErrorCode{30001, "表主键不存在", errors.New("表主键不存在")}
	TableColumnNotExists       = &ErrorCode{30002, "表字段不存在", errors.New("表字段不存在")}
	IndexNotContainsColumn     = &ErrorCode{40001, "索引没有指定字段", errors.New("索引没有指定字段")}
)

func (errCode *ErrorCode) Error() string {
	bytes, err := json.MarshalIndent(errCode, "", "  ") // 使用两个空格作为缩进
	if err != nil {
		// 处理json.MarshalIndent可能发生的错误
		return fmt.Sprintf("{\"error\": \"%v\"}", err)
	}
	return string(bytes)
}
