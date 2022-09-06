package database
import(
      "gorm.io/gorm"
)

type DbLesson struct{
  gorm.Model
  Title       string  `gorm:"index:,size:300;not null"`
  Thumbnail   string  `gorm:"index:,size:120;not null"`
  Min_read    int     `gorm:"index:,size:120;not null"`
  Data        string  `gorm:not null"`
  View        int     `gorm:"index:,size:120;not null"`
}
