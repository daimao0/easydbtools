package enum

import "errors"

type DriverName string

const (
	MySQL      DriverName = "mysql"
	PostgreSQL DriverName = "postgres"
	SQLite     DriverName = "sqlLite"
)

var driverNameMap = map[string]DriverName{
	"mysql":    MySQL,
	"postgres": PostgreSQL,
	"sqlLite":  SQLite,
}

// GetDriverName get enum
func GetDriverName(name string) (DriverName, error) {
	driverName, exists := driverNameMap[name]
	if exists {
		return driverName, nil
	}
	return "", errors.New("driver name not found")
}
