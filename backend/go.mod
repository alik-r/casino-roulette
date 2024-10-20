module github.com/alik-r/casino-roulette/backend

go 1.20

replace github.com/alik-r/casino-roulette/backend/pkg/models => ./pkg/models
replace github.com/alik-r/casino-roulette/backend/pkg/db => ./pkg/db

require (
	gorm.io/driver/sqlite v1.5.6
	gorm.io/gorm v1.25.12
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	golang.org/x/text v0.14.0 // indirect
)
