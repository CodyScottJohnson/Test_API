package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gosuri/uilive"
	"github.com/jinzhu/gorm"
	"github.com/mnmtanish/go-graphiql"
	"github.com/ttacon/chalk"

	"Test_API/graphql"
	"Test_API/models"
)

func Start(db *gorm.DB) error {
	writer := uilive.New()
	writer.Start()
	fmt.Fprintln(writer, chalk.Yellow.Color(chalk.Bold.TextStyle("() Starting Server\n")))
	time.Sleep(time.Millisecond * 5)
	graphql.InititalizeDB(db)
	//db.DropTable(&models.Address{})
	db.AutoMigrate(&models.Address{}, &models.Recruit{}, &models.Tag{})
	//newEmail := models.Address{Address: "1305 south 400 east", City: "Bountiful", State: "Utah", Primary: true, Type: "Work"}
	p := models.Person{}
	db.Where("id = ?", 1).First(&p)
	//db.Model(&p).Association("Addresses").Append(newEmail)
	////////////////
	//Router Set Up
	///////////////
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Cody")
	})
	router.HandleFunc("/Test", func(w http.ResponseWriter, r *http.Request) {
		//vars := mux.Vars(r)
		queryParams := r.URL.Query()
		person := models.Person{}
		_ = json.NewDecoder(r.Body).Decode(&person)
		//  fmt.Fprintf(w,"%#v\n", person)
		fmt.Fprintf(w, "%#v\n", queryParams["cody"])
	}).Methods("POST")
	router.HandleFunc("/graphiql", graphiql.ServeGraphiQL)
	router.HandleFunc("/graphql", graphql.ServeGraphQL())
	fmt.Fprintln(writer, chalk.Blue.Color(chalk.Bold.TextStyle("(*) Server Running\n")))

	//Finish Function
	writer.Stop()
	log.Fatal(http.ListenAndServe(":1234", router))
	return nil
}
