package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ammarfar/mezink-golang-assignment/config"
	"github.com/Ammarfar/mezink-golang-assignment/internal/domain/models"
	database "github.com/Ammarfar/mezink-golang-assignment/pkg/database/gorm"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnvironmentFile(".env")

	db, err := database.NewGorm()
	db.AutoMigrate(
		&models.User{},
		&models.Mark{},
	)

	if err != nil {
		log.Fatalf("Database connection error: %s", err.Error())
	}

	if len(os.Args) < 2 {
		log.Fatalf("Error: Please input an argument to the cli.\n")
	}

	arg := os.Args[1]

	if arg == "up" {
		UpSeedData(db)
		log.Printf("Successfully seeding all data.\n")
	} else {
		log.Fatalf("Error: Only use 'up' or 'down' cli argument for seeding.\n")
	}
}

func UpSeedData(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		var marks = []models.Mark{}
		for j := 0; j < 10; j++ {
			markVal := 1
			if j == 0 {
				markVal = i * 10
			}
			marks = append(marks, models.Mark{
				Mark: &markVal,
			})
		}

		name := fmt.Sprintf("user %d", i)
		userIn := models.User{
			Name:  &name,
			Marks: marks,
		}

		if err := db.Debug().Create(&userIn).Error; err != nil {
			log.Fatalf("\n%s", err.Error())
		}
	}
}
