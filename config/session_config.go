package config

import (
	"os"
	"time"
	"github.com/gofiber/storage/mysql"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
)

func GetSessionStore() *session.Store {

	storage := mysql.New(mysql.Config{
		ConnectionURI: ConnectionString(),
		Reset:         false,
		GCInterval:  time.Duration(sessionExpiration()) * time.Minute,
	})
	store := session.New(session.Config{
        Storage: storage, 
    })
	return store
}

func sessionExpiration() int {
	LoadDotEnv()
	return helpers.ConvertToInt(os.Getenv("APP_SESSION_EXPIRATION"))
}