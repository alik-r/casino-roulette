package main

import (
	"log"
	"net/http"

	"github.com/alik-r/casino-roulette/backend/pkg/api"
	"github.com/alik-r/casino-roulette/backend/pkg/auth"
	"github.com/alik-r/casino-roulette/backend/pkg/db"
	"github.com/alik-r/casino-roulette/backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	db.InitDB()
	auth.InitJWT()

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(corsOptions))

	r.Get("/api/healthcheck", api.Healthcheck)

	r.Post("/api/login", api.Login)
	r.Post("/api/register", api.Register)

	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Get("/api/user", api.GetUser)
		r.Get("/api/leaderboard", api.GetLeaderboard)
		r.Post("/api/roulette", api.PlaceBet)
	})

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
