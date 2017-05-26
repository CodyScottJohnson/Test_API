package main

import (
  "fmt"
  "encoding/json"

  "github.com/ttacon/chalk"

  "Test_API/sqldb"
  "Test_API/models"
  "Test_API/server"
)

func main() {
  fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("-----------")))
  fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("--JFS API--")))
  fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("-----------")))

  db, err := sqldb.Connect()
  defer db.Close()
  err = server.Start()
  if err != nil{
    fmt.Println(chalk.Red.Color( "Unable To Connect to Database:"), err)
    panic(chalk.Red.Color(err.Error()))
  }
//  defer db.Close()

  //fmt.Println(chalk.Blue.Color(chalk.Bold.TextStyle("(*) Connected to Database")))
  //db.DropTable(&Person{}, &Email{}, &Address{})
  //db.AutoMigrate(&Person{}, &Email{}, &Address{})
  s:=models.Person{}
  z :="{ID:1,FirstName:Cody,LastName:Johnsons}"
  json.Unmarshal([]byte(z), &s)
  b2, err := json.Marshal(s)
  fmt.Println(string(b2))
  //db.Create(&p)
  //db.Update(&s)
  p := models.Person{}
  db.Where("id = ?",1).First(&p)
  b, err := json.Marshal(p)
  fmt.Println(string(b))
  //db.Save(&p)

}
