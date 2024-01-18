package config

import (

	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := ConnectionString()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DisconnectDB(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic("Failed to disconnect DB")
	}
	conn.Close()
}

func ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		GetEnv("DB_USER"), 
		GetEnv("DB_PASSWORD"), 
		GetEnv("DB_HOST"),
		GetEnv("DB_PORT"), 
		GetEnv("DB_NAME"),
	)
}