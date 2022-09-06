package model
import(
    "lesson-ms/database"
    // "github.com/mitchellh/mapstructure"
    "fmt"
    "github.com/gofiber/fiber/v2"
)

type Lesson struct{
  Title       string   `json:"Title"`
  Thumbnail   string   `json:"Thumbnail"`
  Min_read    int      `json:"Min_read"`
  Data        string   `json:"Data"`
  View        int      `json:"Views"`
}


func(s *Lesson) getClientInput(c *fiber.Ctx){
  // mapstructure.Decode(c.GetReqHeaders(), &header);//read headers

  if err := c.BodyParser(&s); err != nil {
    fmt.Println("can't body parser line 23. /model/lesson.go")
  }

}


func(s *Lesson) Create(c *fiber.Ctx){
  s.getClientInput(c)
  dbLesson:=new(database.DbLesson)
  dbLesson.Title    =s.Title
  dbLesson.Thumbnail=s.Thumbnail
  dbLesson.Min_read =s.Min_read
  dbLesson.Data     =s.Data
  dbLesson.View     = 0;// mac dinh luc tao chua co view nao

  database.DB.Create(&dbLesson);
}


func(s *Lesson) ViewPost(c *fiber.Ctx){
  id:=c.AllParams()["id"];
  temp:=new(database.DbLesson)
  database.DB.Raw("SELECT * FROM db_lessons WHERE id = ? LIMIT 1", id).Scan(temp);
  increaseView:=temp.View+1;
  database.DB.Raw("UPDATE db_lessons SET View = ? WHERE id = ? LIMIT 1", increaseView,id)
  s.Title     = temp.Title
  s.Thumbnail = temp.Thumbnail
  s.Min_read  = temp.Min_read
  s.Data      = temp.Data
  s.View      = temp.View+1// vi update view
}

func(s *Lesson) Delete(id string){
  database.DB.Delete(&database.DbLesson{}, id)
}
