package expenses

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Expense struct {
	ID    uint
	Value float32
}

func ExpensesRoute(r *gin.Engine, db *gorm.DB) {
	r.GET("/expenses", func(c *gin.Context) {
		var expenses []Expense

		result := db.Model(&Expense{}).Limit(10).Scan(&expenses)

		if err := result.Error; err != nil {
			c.JSON(404, gin.H{
				"message": "Expense not found",
				"error":   err,
			})
		} else {
			c.JSON(200, gin.H{
				"expense": expenses,
			})
		}
	})
}
