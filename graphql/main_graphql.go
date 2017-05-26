package graphql

import(
  "encoding/json"
  "net/http"

  "github.com/graphql-go/graphql"
  "github.com/mnmtanish/go-graphiql"
  "github.com/jinzhu/gorm"
)

var(
  db  *gorm.DB
  schema, err = graphql.NewSchema(graphql.SchemaConfig{
  		Query:    RootQuery,
  		Mutation: RootMutation,
  	})

)
func InititalizeDB(newDB *gorm.DB){
  db = newDB
}
func ServeGraphQL() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		sendError := func(err error) {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		req := &graphiql.Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			sendError(err)
			return
		}

		res := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: req.Query,
		})

		if err := json.NewEncoder(w).Encode(res); err != nil {
			sendError(err)
		}
	}
}
