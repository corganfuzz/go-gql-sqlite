package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Migrate all models
	DB.AutoMigrate(&Project{}, &Skill{}, &Experience{})
}

func SeedDatabase() {
	var count int64
	DB.Model(&Project{}).Count(&count)
	if count > 0 {
		return // Already seeded
	}

	projects := []Project{
		{
			Title:       "Go GraphQL Server",
			Description: "A robust, type-safe GraphQL API server using Go, GORM v2, and SQLite.",
			URL:         "https://github.com/corganfuzz/go-gql-sqlite",
			TechStack:   "Go, GraphQL, SQLite, Docker",
		},
		{
			Title:       "Microservices Operator",
			Description: "Kubernetes Operator for managing microservices at scale.",
			URL:         "https://github.com/example/k8s-operator",
			TechStack:   "Go, Kubernetes, Kubebuilder",
		},
		{
			Title:       "Distributed Cache",
			Description: "High-performance distributed caching system with Raft consensus.",
			URL:         "https://github.com/example/cache-db",
			TechStack:   "Go, gRPC, Protobuf",
		},
	}
	for _, p := range projects {
		DB.Create(&p)
	}

	skills := []Skill{
		{Name: "Go", Category: "Backend", Level: 5},
		{Name: "GraphQL", Category: "API", Level: 4},
		{Name: "Docker", Category: "DevOps", Level: 4},
		{Name: "Kubernetes", Category: "DevOps", Level: 3},
		{Name: "React", Category: "Frontend", Level: 3},
	}
	for _, s := range skills {
		DB.Create(&s)
	}

	experiences := []Experience{
		{
			Company:     "Tech Corp",
			Role:        "Senior Backend Engineer",
			Period:      "2021 - Present",
			Description: "Led migration to Go microservices, improving latency by 40%.",
		},
		{
			Company:     "Startup Inc",
			Role:        "Full Stack Developer",
			Period:      "2019 - 2021",
			Description: "Built MVP using React and Node.js. Scaled to 10k users.",
		},
	}
	for _, e := range experiences {
		DB.Create(&e)
	}

	log.Println("Database seeded with showcase data.")
}
