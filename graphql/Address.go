package graphql

import (
	"github.com/graphql-go/graphql"

	"Test_API/models"
)

var Address = graphql.NewObject(graphql.ObjectConfig{
	Name: "Address",
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
		"Address": &graphql.Field{
			Type: graphql.String,
		},
		"City": &graphql.Field{
			Type: graphql.String,
		},
		"State": &graphql.Field{
			Type: graphql.String,
		},
		"Post": &graphql.Field{
			Type: graphql.String,
		},
		"Type": &graphql.Field{
			Type: graphql.String,
		},
		"Primary": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
