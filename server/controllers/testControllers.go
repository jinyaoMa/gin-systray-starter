package controllers

import (
	"net/http"
	"this/server/services"

	"github.com/gin-gonic/gin"
)

func TestController(c *gin.Context) {
	var testService services.TestService

	err := c.ShouldBindJSON(&testService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tests, err := testService.GetAllTestModels()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "GetAllTestModels() error!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    tests,
	})
}
