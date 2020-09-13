package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"integration-ginkgo-example/internal/config"
	"time"
)

// NewDatabase create a new mysql database with given config
func NewDatabase(conf *config.DatabaseConfig) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	for i := 0; i <= 30; i++ {
		db, err = gorm.Open("mysql", conf.Endpoint)
		if err != nil {
			time.Sleep(500 * time.Millisecond)
		}
	}
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxOpenConns(conf.Pool.MaxOpen)
	db.DB().SetMaxIdleConns(conf.Pool.MaxIdle)
	db.DB().SetConnMaxLifetime(time.Duration(conf.Pool.MaxLifeTime) * time.Second)

	return db, nil
}
