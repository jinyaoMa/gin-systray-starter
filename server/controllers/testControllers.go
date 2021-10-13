package controllers

import (
	"net/http"
	"this/server/services"

	"github.com/gin-gonic/gin"
)

// @Summary TestController
// @Description Test Controller
// @Tags Before Authorization
// @accept plain
// @Produce json
// @Param msg query string false "Msg"
// @Success 200 "{ ok , data }"
// @Failure 400 "binding error"
// @Failure 404 "{ error }"
// @Router /test/ [get]
func TestController(c *gin.Context) {
	var testService services.TestService

	err := c.ShouldBind(&testService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tests, err := testService.GetAllTestModels()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "GetAllTestModels() error!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": tests,
	})
}
