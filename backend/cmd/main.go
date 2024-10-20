package main

import (
	"log"

	"github.com/alik-r/casino-roulette/backend/pkg/db"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Initialize the database
	db.InitDB()
}
