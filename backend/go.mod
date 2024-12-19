module github.com/alik-r/casino-roulette/backend

go 1.22.0

toolchain go1.22.8

replace github.com/alik-r/casino-roulette/backend/pkg/models => ./pkg/models

replace github.com/alik-r/casino-roulette/backend/pkg/db => ./pkg/db

replace github.com/alik-r/casino-roulette/backend/pkg/roulette => ./pkg/roulette

replace github.com/alik-r/casino-roulette/backend/pkg/auth => ./pkg/auth

replace github.com/alik-r/casino-roulette/backend/pkg/middleware => ./pkg/middleware

require (
	golang.org/x/crypto v0.31.0
	gorm.io/gorm v1.25.12
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	golang.org/x/sync v0.10.0 // indirect
)

require (
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-chi/cors v1.2.1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c
	golang.org/x/text v0.21.0 // indirect
	gorm.io/driver/postgres v1.5.9
)
