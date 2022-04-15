package routes

import (
	"dxps.io/go-user-admin-fiber-vue/be/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.RegisterUser)
}
