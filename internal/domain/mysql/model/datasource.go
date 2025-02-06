package model

import "easydbTools/internal/common/enum"

// DataSource datasource model
type DataSource struct {
	// Id datasource id
	Id string

	// DriverName datasource driver name
	DriverName enum.DriverName

	// Name custom name by user
	Name string

	// Address datasource address
	Address string

	// Username datasource username
	Username string

	// Password datasource password
	Password string
}
