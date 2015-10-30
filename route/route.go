package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qwertmax/timeconverter/db"
	// "github.com/qwertmax/timeconverter/model"
)

func Main(c *gin.Context) {
	// Db := GetDB(c)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func GetDB(c *gin.Context) *gorm.DB {
	database := c.MustGet("db").(*db.Database)
	return database.DB
}
