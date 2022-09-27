package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
