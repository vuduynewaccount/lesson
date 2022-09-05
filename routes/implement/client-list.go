package implement
import(
  "github.com/gofiber/fiber/v2"
  "fmt"
)

func List(c *fiber.Ctx) error {
  fmt.Println(c.AllParams()["id"])
  return c.SendString("client viewing list")
}
