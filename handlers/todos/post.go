package todos

import (
	"net/http"

	"github.com/Phazon85/restapp-demo/services/todos"
	"github.com/gin-gonic/gin"
)

// Post godoc
// @Summary      POST todo
// @Description  Creates a new todo
// @Tags         todos
// @Accept	 json
// @Produce      json
// @Param body body PostReq true "Request body."
// @Success      201
// @Failure      500  {object}  nil
// @Router       /todos [post]
func (hand *Handler) Post(c *gin.Context) {
	var req PostReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	//Call Post from todos service.
	if err := hand.service.Post(req.toServiceEntry()); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	//Send Response.
	c.Status(http.StatusCreated)
}

type PostReq struct {
	Name        string `json:"name,omitempty" example:"name"`
	Description string `json:"description,omitempty" example:"This is a description."`
}

func (req PostReq) toServiceEntry() *todos.Entry {
	entry := &todos.Entry{
		Name:        req.Name,
		Description: req.Description,
	}

	return entry
}
