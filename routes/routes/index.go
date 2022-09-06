package routes
import (
    "github.com/gofiber/fiber/v2"
    "lesson-ms/routes/implement"
)

func Path(app *fiber.App){
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  app.Get("/api-client/view/:id",implement.View)
  app.Get("/api-client/list",implement.List)

  app.Get("/api-admin/view/:id",implement.View)
  app.Get("/api-admin/list",implement.List)

  app.Post("/api-admin/create",implement.AdminCreate)
  app.Post("/api-admin/delete/:id",implement.AdminDelete)
  
}
