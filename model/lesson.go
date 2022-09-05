package model
import(
    // "lesson-ms/database"
    // "github.com/mitchellh/mapstructure"
    "fmt"
    "github.com/gofiber/fiber/v2"
)

type Lesson struct{
  Title       string   `json:"Title"`
  Thumbnail   string   `json:"Thumbnail"`
  Min_read    int      `json:"Min_read"`
  Data        string   `json:"Data"`
  Tags        []string `json:"Tags"`
  View        int      `json:"Views"`
}


func(s *Lesson) getClientInput(c *fiber.Ctx){
  fmt.Println("read input from client")
  // mapstructure.Decode(c.GetReqHeaders(), &header);//read headers

  if err := c.BodyParser(&s); err != nil {
    fmt.Println("can't body parser line 24. /model/lesson.go")
  }

}


func(s *Lesson) Create(c *fiber.Ctx){
  s.getClientInput(c)
  s.View = 0;// mac dinh luc tao chua co view nao
  // db.create(s);
}


func(s *Lesson) FindPost(id string){
  
}
