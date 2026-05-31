package main

import (
	"log"

	"bemunair2026/server/config"
	"bemunair2026/server/database"
	"bemunair2026/server/database/entities"
)

func main() {
	cfg := config.Load()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	// Development verification only. Production schema is canonical SQL in server/database/migrations.
	if err := db.AutoMigrate(&entities.User{}, &entities.ContentSubmission{}, &entities.LetterSubmission{}, &entities.MedinfoPJQueue{}, &entities.LetterTemplate{}); err != nil {
		log.Fatal(err)
	}
}
