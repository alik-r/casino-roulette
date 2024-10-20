module github.com/alik-r/casino-roulette/backend

go 1.22.0

toolchain go1.22.8

replace github.com/alik-r/casino-roulette/backend/pkg/models => ./pkg/models

replace github.com/alik-r/casino-roulette/backend/pkg/db => ./pkg/db

replace github.com/alik-r/casino-roulette/backend/pkg/roulette => ./pkg/roulette

require (
	gorm.io/driver/sqlite v1.5.6
	gorm.io/gorm v1.25.12
)

require (
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-chi/cors v1.2.1
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c
	golang.org/x/text v0.19.0 // indirect
)
