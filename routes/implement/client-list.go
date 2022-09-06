package implement
import(
  "github.com/gofiber/fiber/v2"
  "lesson-ms/database"
  "lesson-ms/model"
)

func List(c *fiber.Ctx) error {
  var listLesson []database.DbLesson;
  database.DB.Raw("SELECT * FROM db_lessons").Scan(listLesson);
  resp:=new(model.Resp);
  resp.Constructor(1,"get list thanh cong",listLesson)
  return c.JSON(resp)
}
