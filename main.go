package main

import (
	"github.com/ParampreetWIL/CRUD_Go/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

//	@title			CRUD GO API
//	@version		1.0
//	@description	API for CRUD Operations or a TODO List.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Parampreet Singh Rai
//	@contact.email	parampreets.rai@thewitslab.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:3000
// @BasePath	/
func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Static("/swagger", "./swagger")
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", handlers.GetAllTasksHandler)
	app.Post("/", handlers.AddNewTaskHandler)
	app.Post("/edit", handlers.EditTaskHandler)
	app.Delete("/:id", handlers.DeleteTaskHandler)

	app.Listen(":3000")
}
