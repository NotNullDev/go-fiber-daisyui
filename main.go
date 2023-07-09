package main

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := NewHtmlEngine()
	app := NewApp(engine)

	db := NewDb()

	initHtmlEndpoints(app, db)
	initApiEndpoints(app, db)

	if err := app.Listen(":4444"); err != nil {
		panic(err.Error())
	}
}

func NewHtmlEngine() *html.Engine {
	engine := html.New("./views", ".html")
	engine.Reload(true) // dev only

	engine.AddFunc("getCssAsset", func(name string) (res template.HTML) {
		filepath.Walk("public/assets", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == name {
				res = template.HTML("<link rel=\"stylesheet\" href=\"/" + path + "\">")
			}
			return nil
		})
		return
	})

	return engine
}

func NewApp(engine *html.Engine) *fiber.App {
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	return app

}
