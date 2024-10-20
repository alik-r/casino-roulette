package main

import (
	"log"
	"net/http"

	"github.com/alik-r/casino-roulette/backend/pkg/api"
	"github.com/alik-r/casino-roulette/backend/pkg/db"
	"github.com/go-chi/chi/v5"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	db.InitDB()

	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Post("/user/deposit", api.RegisterOrDeposit)
		r.Post("/roulette/bet", api.PlaceBet)
		r.Get("/leaderboard", api.GetLeaderBoard)
	})

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
