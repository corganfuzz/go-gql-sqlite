package model

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	Name     string
	Category string // e.g., "Frontend", "Backend", "DevOps"
	Level    int    // 1-5
}

var skillType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Skill",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
				Type: graphql.String,
			},
			"level": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

func GetSkills() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(skillType),
		Description: "Get all skills",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var skills []Skill
			err := DB.Find(&skills).Error
			return skills, err
		},
	}
}

func CreateSkill() *graphql.Field {
	return &graphql.Field{
		Type:        skillType,
		Description: "Create a new skill",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"category": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"level": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			skill := Skill{
				Name:     params.Args["name"].(string),
				Category: params.Args["category"].(string),
				Level:    params.Args["level"].(int),
			}
			err := DB.Create(&skill).Error
			return skill, err
		},
	}
}
