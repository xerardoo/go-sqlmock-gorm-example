package models

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() (err error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, password, host, port, database)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		err = fmt.Errorf("connection database: %s", err.Error())
	}

	err = DB.AutoMigrate(&Customer{})
	if err != nil {
		err = fmt.Errorf("database migration: %s", err.Error())
	}
	return
}

func InitMockDB(db *sql.DB) (err error) {
	// https://gorm.io/docs/connecting_to_the_database.html#Existing-database-connection
	DB, err = gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	return
}
