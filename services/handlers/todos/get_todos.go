package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
