package todos

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mailgun/groupcache"
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
	entries, err := hand.service.GetTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	// Send Response.
	c.JSON(http.StatusOK, entries)
}

// Get godoc
// @Summary     GET todos
// @Description Get all todos
// @Tags        todos
// @Produce     json
// @Success     200
// @Failure     404 {object} nil
// @Failure     500 {object} nil
// @Router      /todos [get]
func (hand *Handler) GetByID(c *gin.Context) {
	key := c.Param("key")
	// Call Get from todos service.
	entry, err := hand.service.GetTodoByID(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	// Send Response.
	c.JSON(http.StatusOK, entry)
}

func (hand *Handler) TestGetByID(group *groupcache.Group) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var b []byte
		res := &GetIDRes{}

		err := group.Get(c, c.Param("key"), groupcache.AllocatingByteSliceSink(&b))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		err = json.Unmarshal(b, res)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		// Send Response.
		c.JSON(http.StatusOK, res)
	},
	)
}

type GetIDRes struct {
	ID          string
	Name        string
	Description string
}
