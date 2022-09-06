package database

import(
      "gorm.io/gorm"
      "gorm.io/driver/mysql"
      "fmt"
)

var DB *gorm.DB

var err error

func InitialMigration(){
  var DB_USERNAME string  = env.Get_env("DB_USERNAME")
  var DB_PASSWORD string  = env.Get_env("DB_PASSWORD")
  var DB_HOST     string  = env.Get_env("DB_HOST")
  var DB_PORT     string  = env.Get_env("DB_PORT")
  var DB_DATABASE string  = env.Get_env("DB_DATABASE")


  DBDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT,DB_DATABASE)
  DB,err=gorm.Open(mysql.Open(DBDSN),&gorm.Config{})
  if err!= nil{
    fmt.Println(err.Error())
    panic("can't connect to mysql")
  }

  DB.AutoMigrate(&DbLesson{})
}
