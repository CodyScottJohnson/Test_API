package sqldb

import (
  "fmt"
  "time"
  //"log"
  //"database/sql"
  "github.com/ttacon/chalk"
  "github.com/gosuri/uilive"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB, error){
  writer := uilive.New()
  writer.Start()
  fmt.Fprintf(writer,chalk.Green.Color(chalk.Bold.TextStyle("..Connecting To Database...\n")))
  time.Sleep(time.Millisecond * 5)
  db, err := gorm.Open("mysql", "cody:skiutah4969@tcp(jfsapp.com:3306)/JFS_v2?charset=utf8&parseTime=True&loc=Local")
  if err != nil{
    fmt.Fprintf(writer,chalk.Red.Color( "Unable To Connect to Database:"), err,"\n")
    panic(chalk.Red.Color(err.Error()))

  }
  fmt.Fprintf(writer,chalk.Blue.Color(chalk.Bold.TextStyle("(*) Connected to Database\n")))
  writer.Stop()
  return db, nil

}
