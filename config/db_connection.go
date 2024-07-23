package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := DatabaseURL()
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

func DatabaseURL() string {
	return GetEnv("DATABASE_URL")
}