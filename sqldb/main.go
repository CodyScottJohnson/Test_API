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
  fmt.Fprintln(writer,chalk.Yellow.Color(chalk.Bold.TextStyle("() Connecting To Database")))
  time.Sleep(time.Millisecond * 10)
  db, err := gorm.Open("mysql", "cody:skiutah4969@tcp(jfsapp.com:3306)/JFS_v2?charset=utf8&parseTime=True&loc=Local")
  if err != nil{
    fmt.Fprintln(writer,chalk.Red.Color( "(-)Unable To Connect to Database:"), err)
    panic(chalk.Red.Color(err.Error()))

  }
  fmt.Fprintln(writer,chalk.Blue.Color(chalk.Bold.TextStyle("(*) Connected to Database")))

  writer.Stop()
  return db, nil

}
