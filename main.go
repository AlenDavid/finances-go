package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func before() {
	// load .env file
	err = godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

type Expense struct {
	ID    uint
	Value float32
}

func main() {
	before()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		message := "pong"
		value, exists := c.GetQuery("message")

		if exists {
			message = value
		}

		c.JSON(200, gin.H{
			"message": message,
		})
	})

	r.GET("/expenses", func(c *gin.Context) {
		var expenses []Expense

		result := db.Model(&Expense{}).Limit(10).Scan(&expenses)

		if result.Error != nil {
			c.JSON(404, gin.H{
				"message": "Expense not found",
			})
		}

		c.JSON(200, gin.H{
			"expense": expenses,
		})
	})

	r.Run() // listen and server on 0.0.0.0:8080
}
