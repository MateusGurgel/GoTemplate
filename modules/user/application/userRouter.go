package application

import (
	"github.com/gofiber/fiber/v2"
	"goTemplate/modules/user/application/controller/createUserController"
)

func SetupUserRouter(app *fiber.App) {
	userRouter := app.Group("/users")
	userRouter.Post("/", createUserController.Handle)
}
