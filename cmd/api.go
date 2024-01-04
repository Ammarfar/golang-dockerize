package main

import (
	"log"

	"github.com/Ammarfar/mezink-golang-assignment/config"
	database "github.com/Ammarfar/mezink-golang-assignment/pkg/database/gorm"
	"github.com/Ammarfar/mezink-golang-assignment/server"
)

func main() {
	log.Println("Starting api server")

	config.LoadEnvironmentFile(".env")

	db, err := database.NewGorm()

	if err != nil {
		log.Fatalf("db init: %s", err)
	}

	s := server.NewServer(db)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
