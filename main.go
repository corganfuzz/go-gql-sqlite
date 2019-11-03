package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/corganfuzz/go-gql-sqlite/pkg/model"
	"github.com/graphql-go/graphql"
)

// // Tutorial comment
// type Tutorial struct {
// 	ID       int
// 	Title    string
// 	Author   Author
// 	Comments []Comment
// }

// // Author comment
// type Author struct {
// 	Name      string
// 	Tutorials []int
// }

// // Comment here
// type Comment struct {
// 	Body string
// }

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Println("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

var agregateSchema = graphql.Fields{
	"tutorial": model.SingleTutorialSchema(),
	"list":     model.ListTutorialSchema(),
}

var agregateMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": model.CreateTutorialMutation(),
	},
})

// var authorType = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "Author",
// 		Fields: graphql.Fields{
// 			"Name": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"Tutorials": &graphql.Field{
// 				Type: graphql.NewList(graphql.Int),
// 			},
// 		},
// 	},
// )

// var commentType = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "Comment",
// 		Fields: graphql.Fields{
// 			"body": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 		},
// 	},
// )

// var tutorialType = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "Tutorial",
// 		Fields: graphql.Fields{
// 			"id": &graphql.Field{
// 				Type: graphql.Int,
// 			},
// 			"title": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"author": &graphql.Field{
// 				Type: authorType,
// 			},
// 			"comments": &graphql.Field{
// 				Type: graphql.NewList(commentType),
// 			},
// 		},
// 	},
// )

// var mutationType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Mutation",
// 	Fields: graphql.Fields{
// 		"create": &graphql.Field{
// 			Type:        tutorialType,
// 			Description: "Create a new Tutorial",
// 			Args: graphql.FieldConfigArgument{
// 				"id": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.Int),
// 				},
// 				"title": &graphql.ArgumentConfig{
// 					Type: graphql.NewNonNull(graphql.String),
// 				},
// 			},
// 			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 				db, err := sql.Open("sqlite3", "./tutorials.db")
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				defer db.Close()

// 				stmt, err := db.Prepare("INSERT INTO tutorials VALUES (?, ?)")
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				defer stmt.Close()

// 				_, err = stmt.Exec(params.Args["id"].(int), params.Args["title"].(string))
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 				var tutorial Tutorial
// 				err = db.QueryRow("SELECT * FROM tutorials where ID = ?", params.Args["id"].(int)).Scan(&tutorial.ID, &tutorial.Title)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 				return tutorial, nil
// 			},
// 		},
// 	},
// })

func main() {
	// db, err := sql.Open("sqlite3", "./tutorials.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// // Schema
	// fields := graphql.Fields{
	// 	"tutorial": &graphql.Field{
	// 		Type:        tutorialType,
	// 		Description: "Get Tutorial By ID",
	// 		Args: graphql.FieldConfigArgument{
	// 			"id": &graphql.ArgumentConfig{
	// 				Type: graphql.Int,
	// 			},
	// 		},
	// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// 			id, ok := p.Args["id"].(int)
	// 			if ok {
	// 				var tutorial Tutorial
	// 				err := db.QueryRow("SELECT * FROM tutorials where ID = ?", id).Scan(&tutorial.ID, &tutorial.Title)
	// 				if err != nil {
	// 					fmt.Println(err)
	// 				}
	// 				return tutorial, nil
	// 			}
	// 			return nil, nil
	// 		},
	// 	},
	// 	"list": &graphql.Field{
	// 		Type:        graphql.NewList(tutorialType),
	// 		Description: "Get Tutorial List",
	// 		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
	// 			db, err := sql.Open("sqlite3", "./tutorials.db")
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 			defer db.Close()
	// 			// perform a db.Query insert
	// 			var tutorials []Tutorial
	// 			results, err := db.Query("SELECT * FROM tutorials")
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			}
	// 			for results.Next() {
	// 				var tutorial Tutorial
	// 				err = results.Scan(&tutorial.ID, &tutorial.Title)
	// 				if err != nil {
	// 					fmt.Println(err)
	// 				}
	// 				log.Println(tutorial)
	// 				tutorials = append(tutorials, tutorial)
	// 			}
	// 			return tutorials, nil
	// 		},
	// 	},
	// }

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: agregateSchema}
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    graphql.NewObject(rootQuery),
			Mutation: agregateMutations,
		},
	)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		mutation {
			create(id: 3, title: "Acid") {
				id
				title
			}
		}
	`
	// params := graphql.Params{Schema: schema, RequestString: query}
	// r := graphql.Do(params)
	// if len(r.Errors) > 0 {
	// 	log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	// }
	// rJSON, _ := json.Marshal(r)

	result := executeQuery(query, schema)
	rJSON, _ := json.Marshal(result)
	fmt.Printf("%s \n", rJSON)

}
