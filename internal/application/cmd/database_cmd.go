package cmd

// DatabaseCreateCmd app create database
type DatabaseCreateCmd struct {
	//DataSourceId
	DataSourceId string
	// Name
	Name string
}

// DatabaseDropCmd app drop database
type DatabaseDropCmd struct {
	//DataSourceId
	DataSourceId string
	// Name
	Name string
}
