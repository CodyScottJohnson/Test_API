package graphql

import(
  "github.com/graphql-go/graphql"

    "Test_API/models"
)

var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.String,
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			     return p.Source.(models.Person).ID, nil
			},
		},
    "CreatedAt": &graphql.Field{
			Type: graphql.String,
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			     return p.Source.(models.Person).CreatedAt, nil
			},
		},
    "UpdatedAt": &graphql.Field{
			Type: graphql.String,
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			     return p.Source.(models.Person).UpdatedAt, nil
			},
		},
    "FirstName": &graphql.Field{
			Type: graphql.String,
		},
    "LastName": &graphql.Field{
      Type: graphql.String,
    },
	},

})
