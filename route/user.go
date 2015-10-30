package route

import (
	"github.com/gin-gonic/gin"
	"github.com/qwertmax/timeconverter/model"
)

func UsersList(c *gin.Context) {
	Db := GetDB(c)

	var users []model.User
	err := Db.Find(&users).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

func UserCreate(c *gin.Context) {
	Db := GetDB(c)

	var userNew model.User
	err := c.BindJSON(&userNew)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"Message": err.Error(),
		})
		return
	}

	err = Db.Save(&userNew).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(201, userNew)
}

func UserGet(c *gin.Context) {
	Db := GetDB(c)
	id := c.Param("id")

	var user model.User
	err := Db.First(&user, id).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func UserUpdate(c *gin.Context) {
	Db := GetDB(c)
	id := c.Param("id")

	var userNew model.User
	err := c.BindJSON(&userNew)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"Message": err.Error(),
		})
		return
	}

	var user model.User
	err = Db.First(&user, id).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"Message": err.Error(),
		})
		return
	}

	userNew.ID = user.ID
	err = Db.Save(&userNew).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(200, userNew)
}

func UserDelete(c *gin.Context) {
	Db := GetDB(c)
	id := c.Param("id")

	err := Db.Delete(&model.User{}, id).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "ok",
		"Message": "User was deleted successfully",
	})
}
