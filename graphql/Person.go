package graphql

import (
	"github.com/graphql-go/graphql"

	"Test_API/models"
)

var Person = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
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
		"Primary_Email": &graphql.Field{
			Type: graphql.String,
		},
		"Primary_Address": &graphql.Field{
			Type: graphql.String,
		},
		"Primary_Phone": &graphql.Field{
			Type: graphql.String,
		},
		"Emails": &graphql.Field{
			Type: graphql.NewList(Email),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(models.Person)
				emails := []models.Email{}
				db.Model(&user).Association("Emails").Find(&emails)
				return emails, nil
			},
		},
		"Addresses": &graphql.Field{
			Type: graphql.NewList(Address),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(models.Person)
				addresses := []models.Address{}
				db.Model(&user).Association("Emails").Find(&addresses)
				return addresses, nil
			},
		},
	},
})
