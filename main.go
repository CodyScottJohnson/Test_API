package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"

  "github.com/ttacon/chalk"
  "github.com/gorilla/mux"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "Test_API/sqldb"
  "Test_API/models"
)

func main() {
  fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("-------------------")))
  fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("--Starting Server--")))
  fmt.Println(chalk.Green.Color(chalk.Bold.TextStyle("-------------------")))

  db, err := sqldb.Connect()
  defer db.Close()
//  if err != nil{
//    fmt.Println(chalk.Red.Color( "Unable To Connect to Database:"), err)
  //  panic(chalk.Red.Color(err.Error()))
  //}
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
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello Cody", db, err)
  })
  router.HandleFunc("/Test", func(w http.ResponseWriter, r *http.Request){
    person := models.Person{}
    _ = json.NewDecoder(r.Body).Decode(&person)
    fmt.Fprintf(w,"%#v\n", person)
  }).Methods("POST")
  log.Fatal(http.ListenAndServe(":1234", router))
}
