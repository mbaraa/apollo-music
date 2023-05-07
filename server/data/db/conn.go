package db

import (
	"fmt"

	"github.com/mbaraa/apollo-music/config/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var instance *gorm.DB = nil

// GetDBConnector returns a singleton mysql connection instance to the application's DB
func GetDBConnector() *gorm.DB {
	return getDBConnector("apollo_music")
}

// GetTestDBConnector returns a singleton mysql connection instance to the application's test DB
func GetTestDBConnector() *gorm.DB {
	return getDBConnector("apollo_music_test").Debug()
}

// getDBConnector returns a singleton mysql connection instance
func getDBConnector(dbName string) *gorm.DB {
	if instance == nil {
		var err error
		createDBDsn := fmt.Sprintf("%s:%s@tcp(%s)/", env.DBUser(), env.DBPassword(), env.DBHost())
		database, err := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{})

		_ = database.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + ";")

		instance, err = gorm.Open(mysql.New(mysql.Config{
			DriverName: "mysql",
			DSN:        fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", env.DBUser(), env.DBPassword(), env.DBHost(), dbName),
		}))
		if err != nil {
			panic(err)
		}
	}
	return instance
}
