package server

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/ttacon/chalk"
  "github.com/gosuri/uilive"
  "github.com/gorilla/mux"
  "Test_API/models"
)

func Start()(error){
writer := uilive.New()
writer.Start()
fmt.Fprintf(writer,chalk.Yellow.Color(chalk.Bold.TextStyle("() Starting Server\n")))
router := mux.NewRouter().StrictSlash(true)
router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello Cody")
})
router.HandleFunc("/Test", func(w http.ResponseWriter, r *http.Request){
  person := models.Person{}
  _ = json.NewDecoder(r.Body).Decode(&person)
  fmt.Fprintf(w,"%#v\n", person)
}).Methods("POST")
fmt.Fprintf(writer,chalk.Blue.Color(chalk.Bold.TextStyle("(*) Server Running\n")))
writer.Stop()
log.Fatal(http.ListenAndServe(":1234", router))
return nil
}
