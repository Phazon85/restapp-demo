package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delete godoc
// @Summary     DELETE todo
// @Description Deletes a todo by id
// @Tags        todos
// @Produce     json
// @Param       body body PostReq true "Request body."
// @Success     200
// @Failure     500 {object} nil
// @Router      /todos/:id [delete]
func (hand *Handler) Delete(c *gin.Context) {
	// Call Delete from todos service
	if err := hand.service.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// Send Response.
	c.Status(http.StatusOK)
}
