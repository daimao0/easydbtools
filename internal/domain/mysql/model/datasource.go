package model

import "easydbTools/internal/common/enum"

type DataSource struct {
	Id string

	DriverName enum.DriverName

	Name string

	Address string

	Username string

	Password string
}
