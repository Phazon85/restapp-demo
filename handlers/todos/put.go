package todos

import (
	"net/http"

	"github.com/Phazon85/restapp-demo/services/todos"
	"github.com/gin-gonic/gin"
)

// Put godoc
// @Summary      Put todo
// @Description  Updates the name or description of a particular todo
// @Tags         todos
// @Produce      json
// @Param body body PutReq true "Request body."
// @Success      200
// @Failure      500  {object}  nil
// @Router       /todos/:id [delete]
func (hand *Handler) Put(c *gin.Context) {
	var req PutReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	//Call Put from todos service
	if err := hand.service.Put(req.toServiceEntry(c.Param("id"))); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	//Send Response.
	c.Status(http.StatusOK)
}

type PutReq struct {
	ID          string
	Name        string
	Description string
}

func (req PutReq) toServiceEntry(id string) *todos.Entry {
	entry := &todos.Entry{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	return entry
}
