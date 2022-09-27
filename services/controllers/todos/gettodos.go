package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTodos godoc
// @Summary      Get all todos
// @Description  Get all todos
// @Tags         todos
// @Produce      json
// @Success      200
// @Failure      404  {object}  nil
// @Failure      500  {object}  nil
// @Router       /api/v1/todos [get]
func (ctrl *Controller) GetTodos(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "test",
	})
}
