package taskcontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/marwenbhriz/ga-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var tasks []models.Task

	models.DB.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})

}

func Show(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := models.DB.First(&task, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Show tasks failed"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func Create(c *gin.Context) {

	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		log.Fatal(err)
		return
	}

	models.DB.Create(&task)
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func Update(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&task).Where("id = ?", id).Updates(&task).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Update request failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task details updated"})

}

func Delete(c *gin.Context) {

	var task models.Task

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&task, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Task detailed failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}
