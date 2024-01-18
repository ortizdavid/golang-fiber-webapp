package config

import (
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
	return GetEnv("APP_HOST")+":"+GetEnv("APP_PORT")
}

func ItemsPerPage() int {
	return helpers.ConvertToInt(GetEnv("APP_ITEMS_PER_PAGE"))
}

func UploadImagePath() string {
	return GetEnv("UPLOAD_IMAGE_PATH")
}

func UploadDocumentPath() string {
	return GetEnv("UPLOAD_DOCUMENT_PATH")
}
