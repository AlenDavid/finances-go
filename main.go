package main

import (
	"os"

	"davidalen.dev/finances/expenses"
	"davidalen.dev/finances/ping"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func before() {
	// load .env file
	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

func main() {
	before()

	r := gin.Default()

	ping.PingRoute(r)
	expenses.ExpensesRoute(r, db)

	r.Run() // listen and server on 0.0.0.0:8080
}
