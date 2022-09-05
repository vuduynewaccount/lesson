package implement
import(
  "github.com/gofiber/fiber/v2"
  "fmt"
)

func AdminDelete(c *fiber.Ctx) error {
  fmt.Println(c.AllParams()["id"])
  return c.SendString("admin delete post")
}
