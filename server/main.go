package main

import (
	"log"
	"os"

	controller "github.com/Aritra640/AnonymousChatApp/server/Controller"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

  if err := godotenv.Load(); err != nil {
    log.Println("Error: cannot find .env file")
    panic("Error in godotenv!")
  }
  port := os.Getenv("PORT")

  e := echo.New()

  controller.GroupRoutes(e)

  e.Logger.Fatal(e.Start(port))
}
