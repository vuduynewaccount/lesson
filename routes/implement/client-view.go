package implement
import(
  "github.com/gofiber/fiber/v2"
  "lesson-ms/model"
  "fmt"
)

func View(c *fiber.Ctx) error {
  lesson:=new(model.Lesson);
  lesson.ViewPost(c);// truyen c vao vi body input nam o trong c
  resp:=new(model.Resp);
  resp.Constructor(1,"view thanh cong",lesson)
  return c.JSON(resp)
}
