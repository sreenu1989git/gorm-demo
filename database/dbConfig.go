package database

import (
	"example/gorm/database/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var databaseInstance *gorm.DB

func Init() *gorm.DB {
	var err error

	databaseInstance, err = getDatabaseConnection()
	if err != nil {
		log.Fatalf("Could not connect database: %v", err)
	}

	err = createTable()
	if err != nil {
		log.Fatalf("Could not create table: %v", err)
	}
	return databaseInstance
}

func getDatabaseConnection() (*gorm.DB, error) {
	dsn := "root:12345678@tcp(localhost:3306)/practice?parseTime=true&loc=Local"

	databaseConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal(err)
	}
	return databaseConnection, nil
}

func createTable() error {
	err := databaseInstance.AutoMigrate(&models.Countries{})

	if err != nil {
		return err
	}
	return nil
}
