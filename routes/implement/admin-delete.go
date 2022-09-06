package implement
import(
  "github.com/gofiber/fiber/v2"
  "lesson-ms/model"
)

func AdminDelete(c *fiber.Ctx) error {
  lesson:=new(model.Lesson);


  lesson.Delete(c.AllParams()["id"])// truyen id vao

  resp:=new(model.Resp);
  resp.Constructor(1,"delete successly","")

  return c.JSON(resp)
}
