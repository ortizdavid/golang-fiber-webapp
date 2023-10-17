package config

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
)

func ConfigStaticFiles(app *fiber.App) {
	app.Static("/", "./static")
}

func GetTemplateEngine() *html.Engine {
	engine := html.New("./templates", ".html")
	return engine
}

func ListenAddr() string {
	LoadDotEnv()
	return os.Getenv("APP_HOST")+":"+os.Getenv("APP_PORT")
}

func ItemsPerPage() int {
	LoadDotEnv()
	return helpers.ConvertToInt(os.Getenv("APP_ITEMS_PER_PAGE"))
}

func UploadImagePath() string {
	LoadDotEnv()
	return os.Getenv("UPLOAD_IMAGE_PATH")
}

func UploadDocumentPath() string {
	LoadDotEnv()
	return os.Getenv("UPLOAD_DOCUMENT_PATH")
}
