package implement
import(
  "github.com/gofiber/fiber/v2"
  "lesson-ms/model"
  "fmt"
)

func View(c *fiber.Ctx) error {
  fmt.Println(c.AllParams()["id"])
  resp:=new(model.Resp);
  resp.Constructor(1,"view thanh cong","data")
  return c.JSON(resp)
}
