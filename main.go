package main

import (
  "fmt"

  "github.com/ttacon/chalk"

  "Test_API/sqldb"
  "Test_API/server"
)

func main() {
  fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("-----------")))
  fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("--JFS API--")))
  fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("-----------")))

  db, err := sqldb.Connect()
  defer db.Close()
  err = server.Start(db)
  if err != nil{
    fmt.Println(chalk.Red.Color( "Unable To Connect to Database:"), err)
    panic(chalk.Red.Color(err.Error()))
  }
}
