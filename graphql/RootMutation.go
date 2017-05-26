package graphql

import(
  "github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"setMessage": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"msg": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				msg := "John"
				return msg, nil
			},
		},
	},
})
