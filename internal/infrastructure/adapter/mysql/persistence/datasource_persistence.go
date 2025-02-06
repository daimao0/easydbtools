package persistence

import (
	"context"
	"database/sql"
	"easydbTools/internal/common/error_code"
	"easydbTools/internal/domain/mysql/model"
	"fmt"
	"log"
	"sync"
	"time"
)

// DataSourceRepositoryImpl implements the DataSourceRepository interface
type DataSourceRepositoryImpl struct {
	// mu read-write lock
	mu    sync.RWMutex
	dbMap map[string]*sql.DB
}

var (
	instance *DataSourceRepositoryImpl
	once     sync.Once
)

func NewDataSourceRepositoryImpl() *DataSourceRepositoryImpl {
	once.Do(func() {
		instance = &DataSourceRepositoryImpl{
			dbMap: make(map[string]*sql.DB),
		}
	})
	return instance
}

// TestConnect to the database
func (repo *DataSourceRepositoryImpl) TestConnect(dataSource model.DataSource) error {
	driverName := string(dataSource.DriverName)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/", dataSource.Username, dataSource.Password, dataSource.Address)
	open, err := sql.Open(driverName, dataSourceName)
	if open == nil || err != nil {
		return error_code.DataSourceConnectError
	}
	defer func(open *sql.DB) {
		err := open.Close()
		if err != nil {
			log.Fatalf("close db error: %v", err)
		}
	}(open)
	// ping context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return open.PingContext(ctx)
}

// Connect to the database and cache in memory
func (repo *DataSourceRepositoryImpl) Connect(dataSource model.DataSource) (*sql.DB, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// check whether the database has been connected
	if db, exists := repo.dbMap[dataSource.Id]; exists {
		return db, nil
	}
	// connect to the database
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/", dataSource.Username, dataSource.Password, dataSource.Address)
	db, err := sql.Open(string(dataSource.DriverName), dataSourceName)
	if err != nil {
		return nil, err
	}
	// set connections pool parameters
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	// ping context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	// cache in memory
	repo.dbMap[dataSource.Id] = db
	return db, nil
}

func (repo *DataSourceRepositoryImpl) ConnectById(dataSourceId string) *sql.DB {
	if db, exists := repo.dbMap[dataSourceId]; exists {
		return db
	}
	return nil
}
