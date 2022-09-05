package implement
import(
  "github.com/gofiber/fiber/v2"
  "lesson-ms/model"
)

func AdminCreate(c *fiber.Ctx) error {

  lesson:=new(model.Lesson);


  lesson.Create(c)// truyen c vao vi body input nam o trong c

  resp:=new(model.Resp);
  resp.Constructor(1,"create successly",lesson)

  return c.JSON("admin create list")
}
