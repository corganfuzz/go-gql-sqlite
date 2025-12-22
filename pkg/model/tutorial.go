package model

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Tutorial struct {
	gorm.Model
	ID       int
	Title    string
	Author   Author
	Comments []Comment
}

var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: authorType,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
			},
		},
	},
)

var DB *gorm.DB

func SetupDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "tutorials.db")
	if err != nil {
		log.Fatal(err)
	}
	DB.LogMode(true)
	DB.AutoMigrate(&Tutorial{}, &Comment{}, &Author{})
}

func SingleTutorialSchema() *graphql.Field {
	return &graphql.Field{
		Type:        tutorialType,
		Description: "Get Tutorial By ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var tutorial Tutorial
			// Use shared DB
			DB.Preload("Author").Preload("Comments").First(&tutorial, params.Args["id"].(int))
			return tutorial, nil
		},
	}
}

func ListTutorialSchema() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(tutorialType),
		Description: "Get Tutorial List",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var tutorials []Tutorial
			// Use shared DB
			DB.Preload("Author").Preload("Comments").Find(&tutorials)
			return tutorials, nil
		},
	}

}

func CreateTutorialMutation() *graphql.Field {
	return &graphql.Field{
		Type:        tutorialType,
		Description: "Create new tutorial",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			tutorial := Tutorial{ID: params.Args["id"].(int), Title: params.Args["title"].(string)}
			// Use shared DB
			DB.Create(&tutorial)
			return tutorial, nil
		},
	}
}
