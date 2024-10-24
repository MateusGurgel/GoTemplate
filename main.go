package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"goTemplate/infra/database"
	"goTemplate/modules/user/application"
	"log"
)

func main() {
	fmt.Println("The app is starting")
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	database.Database.InitDatabase()

	app := fiber.New()
	application.SetupUserRouter(app)
	log.Fatal(app.Listen(":8000"))
}
