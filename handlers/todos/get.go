package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get godoc
// @Summary     GET todos
// @Description Get all todos
// @Tags        todos
// @Produce     json
// @Success     200
// @Failure     404 {object} nil
// @Failure     500 {object} nil
// @Router      /todos [get]
func (hand *Handler) Get(c *gin.Context) {
	// Call Get from todos service.
	entries, err := hand.service.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// Send Response.
	c.JSON(http.StatusOK, entries)
}
