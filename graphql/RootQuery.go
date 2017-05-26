package graphql

import(
  "fmt"
  "encoding/json"
  "github.com/graphql-go/graphql"
  "Test_API/models"

)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getMessage": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				msg := "Cody"
				return msg, nil
			},
		},
    "me": &graphql.Field{
      Type:User,
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        user := models.Person{}
        db.Where("id = ?",1).First(&user)
        b, _ := json.Marshal(user.ID)
        fmt.Println(string(b))
				return user, nil
			},
    },
	},

})
