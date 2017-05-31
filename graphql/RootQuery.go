package graphql

import (
	"Test_API/models"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
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
			Type: User,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := models.Person{}
				db.Where("id = ?", 1).First(&user)
				b, _ := json.Marshal(user.ID)
				fmt.Println(string(b))
				return user, nil
			},
		},
		"recruits": &graphql.Field{
			Type: graphql.NewList(Recruit),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				recruit := []models.Recruit{}
				db.Preload("Person").Where("archived = ?", false).Find(&recruit)
				return recruit, nil
			},
		},
	},
})
