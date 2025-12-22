package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/corganfuzz/go-gql-sqlite/pkg/model"
	"github.com/graphql-go/graphql"
)

var aggregateSchema = graphql.Fields{
	"tutorial": model.SingleTutorialSchema(),
	"list":     model.ListTutorialSchema(),
}

var aggregateMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": model.CreateTutorialMutation(),
	},
})

func main() {

	// Initialize DB

	model.SetupDB()
	defer model.DB.Close()

	// Setup Schema

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: aggregateSchema}
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    graphql.NewObject(rootQuery),
			Mutation: aggregateMutations,
		},
	)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// HTTP Handler
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var p struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  p.Query,
			OperationName:  p.OperationName,
			VariableValues: p.Variables,
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	// Serve GraphiQL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	})

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
