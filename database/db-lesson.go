package database
import(
      "gorm.io/gorm"
)

type DbLesson struct{
  gorm.Model
  Title       string
  Thumbnail   string
  Min_read    int
  Data        string
  View        int      
}
