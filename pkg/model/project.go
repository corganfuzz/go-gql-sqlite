package model

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title       string
	Description string
	URL         string
	TechStack   string // Space or comma separated
}

var projectType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Project",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
			"techStack": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func GetProjects() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(projectType),
		Description: "Get all projects",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var projects []Project
			err := DB.Find(&projects).Error
			return projects, err
		},
	}
}

func CreateProject() *graphql.Field {
	return &graphql.Field{
		Type:        projectType,
		Description: "Create a new project",
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"url": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"techStack": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			project := Project{
				Title:       params.Args["title"].(string),
				Description: params.Args["description"].(string),
				URL:         params.Args["url"].(string),
				TechStack:   params.Args["techStack"].(string),
			}
			err := DB.Create(&project).Error
			return project, err
		},
	}
}
