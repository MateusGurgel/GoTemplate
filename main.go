package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"goTemplate/infra/database"
	"goTemplate/modules/user/services/createUser"
)

func main() {
	godotenv.Load()

	fmt.Println("The app is starting")
	database.Database.InitDatabase()

	createUser.CreatUser(&createUser.DTO{Email: "mateusgurgel15@gmail.com", Password: "123"})

}
