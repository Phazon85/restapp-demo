package todos

import (
	"net/http"

	"github.com/Phazon85/restapp-demo/services/todos"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *todos.Service
}

func New(repo *todos.Service) *Handler {
	return &Handler{
		service: repo,
	}
}

// GetTodos godoc
// @Summary      Get all todos
// @Description  Get all todos
// @Tags         todos
// @Produce      json
// @Success      200
// @Failure      404  {object}  nil
// @Failure      500  {object}  nil
// @Router       /todos [get]
func (hand *Handler) GetTodos(c *gin.Context) {
	// Call GetTodos from Repository.
	entries, err := hand.service.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	//Send Response.
	c.JSON(http.StatusOK, entries)
}
