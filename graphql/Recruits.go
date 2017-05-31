package graphql

import (
	"github.com/graphql-go/graphql"

	"Test_API/models"
)

var Recruit = graphql.NewObject(graphql.ObjectConfig{
	Name: "Recruit",
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
		"Person": &graphql.Field{
			Type: Person,
		},
	},
})
