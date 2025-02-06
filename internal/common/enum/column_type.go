package enum

import "errors"

type ColumnType string

const (
	INT       ColumnType = "int"
	BIGINT    ColumnType = "bigint"
	DECIMAL   ColumnType = "decimal"
	VARCHAR   ColumnType = "varchar"
	TEXT      ColumnType = "text"
	DATETIME  ColumnType = "datetime"
	TIMESTAMP ColumnType = "timestamp"
)

var columnTypeMap = map[string]ColumnType{
	"int":       INT,
	"bigint":    BIGINT,
	"decimal":   DECIMAL,
	"varchar":   VARCHAR,
	"text":      TEXT,
	"DATETIME":  DATETIME,
	"timestamp": TIMESTAMP,
}

// GetColumnType get enum
func GetColumnType(name string) (ColumnType, error) {
	columnType, exists := columnTypeMap[name]
	if exists {
		return columnType, nil
	}
	return "", errors.New("column type not found")
}
