package config

import (
	"time"
	"github.com/gofiber/storage/mysql"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
)

func GetSessionStore() *session.Store {
	storage := mysql.New(mysql.Config{
		ConnectionURI: DatabaseURL(),
		Reset:         false,
		GCInterval:  time.Duration(sessionExpiration()) * time.Minute,
	})
	store := session.New(session.Config{
        Storage: storage, 
    })
	return store
}

func sessionExpiration() int {
	return helpers.ConvertToInt(GetEnv("APP_SESSION_EXPIRATION"))
}