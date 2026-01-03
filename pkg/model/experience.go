package model

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	Company     string
	Role        string
	Period      string
	Description string // Accomplishments/Details
}

var experienceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Experience",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"company": &graphql.Field{
				Type: graphql.String,
			},
			"role": &graphql.Field{
				Type: graphql.String,
			},
			"period": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func GetExperiences() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(experienceType),
		Description: "Get all experiences",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var experiences []Experience
			err := DB.Find(&experiences).Error
			return experiences, err
		},
	}
}

func CreateExperience() *graphql.Field {
	return &graphql.Field{
		Type:        experienceType,
		Description: "Create a new experience entry",
		Args: graphql.FieldConfigArgument{
			"company": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"role": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"period": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			exp := Experience{
				Company:     params.Args["company"].(string),
				Role:        params.Args["role"].(string),
				Period:      params.Args["period"].(string),
				Description: params.Args["description"].(string),
			}
			err := DB.Create(&exp).Error
			return exp, err
		},
	}
}
