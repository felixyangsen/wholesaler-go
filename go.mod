// +heroku goVersion go1.13
module myapp

go 1.13

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.8.0
	github.com/rs/cors v1.7.0
	github.com/vektah/dataloaden v0.3.0 // indirect
	github.com/vektah/gqlparser v1.3.1
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.20.12
)
