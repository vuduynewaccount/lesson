package main

import (
    "github.com/gofiber/fiber/v2"
    "lesson-ms/routes/routes"
    "lesson-ms/database"
    "fmt"
    "github.com/joho/godotenv"
)



func main() {
  app := fiber.New()

  app.Get("/api/health", func(c *fiber.Ctx)error{
    return c.SendString("test health")
  })
  // check enviroment variable
  err := godotenv.Load("app.env")
	if err != nil {
		fmt.Println("Fiber can't load env  ", err)
	}
  // connect to database
  database.InitialMigration()

  routes.Path(app)
  if err := app.Listen(":3000"); err != nil {
		fmt.Println("Fiber server got error ", err)
	}
}
